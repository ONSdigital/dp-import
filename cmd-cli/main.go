package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"

	"github.com/ONSdigital/dp-dataset-api/models"
	imodel "github.com/ONSdigital/dp-import-api/models"
	"github.com/ONSdigital/dp-import/cmd-cli/config"
	"github.com/ONSdigital/dp-import/cmd-cli/requests"
	"github.com/ONSdigital/go-ns/log"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const DATASET_ID = "931a8a2a-0dc8-42b6-a884-7b6054ed3b68"

var (
	logReq  bool
	logResp bool
	client  = http.Client{}
	cfg     config.Model
)

func main() {
	log.HumanReadable = true
	cmd := flag.String("cmd", "", "")
	logRequests := flag.Bool("logReq", true, "set true to log out the JSON requests")
	logResponses := flag.Bool("logResp", true, "set true to log out the JSON responses")
	flag.Parse()

	logReq = *logRequests
	logResp = *logResponses

	cfg = config.Load()
	requests.Cfg = cfg
	requests.LogRequests = logReq

	command := strings.ToLower(*cmd)

	switch command {
	case "clean":
		DropDatabases()
	case "clean-import":
		DropDatabases()
		ImportDataset()
	case "get-dataset":
		body := makeRequest(requests.GetDataset(DATASET_ID), http.StatusOK)
		var dataset models.Dataset
		displayResponse(body, dataset, "GetDataset")
	default:
		log.Error(errors.New("unknown cmd value"), log.Data{"cmd": command})
	}
}

func ImportDataset() {
	log.Info("automating dataset import", nil)
	log.Info("beginning import...", nil)

	var body []byte
	var dataset models.Dataset

	// Create dataset
	body = makeRequest(requests.CreateDataset(DATASET_ID), http.StatusCreated)
	displayResponse(body, &dataset, "CreateDataset")

	// Upload file
	var job imodel.Job
	body = makeRequest(requests.UploadFile(), http.StatusCreated)
	displayResponse(body, &job, "UploadFile")

	// capture the instance created by the job
	fmt.Println(job.Links.Instances)
	jobInstanceID := job.Links.Instances[0].ID

	// Submit job
	body = makeRequest(requests.Submit(job.ID), http.StatusOK)
	displayResponse(body, &job, "Submit")

	// Get instances
	var instances models.InstanceResults
	body = makeRequest(requests.GetInstances(), http.StatusOK)
	displayResponse(body, &instances, "GetInstances")

	var instance models.Instance
	for _, i := range instances.Items {
		if i.InstanceID == jobInstanceID {
			instance = i
			break
		}
	}

	log.Info("waiting for instance to ready...", nil)
	time.Sleep(time.Second * 5)

	for {

		body = makeRequest(requests.GetInstance(instance.InstanceID), http.StatusOK)
		json.Unmarshal(body, &instance)
		if instance.State == "completed" {
			break
		}
		time.Sleep(time.Second * 5)
	}

	body = makeRequest(requests.UpdateInstance(instance.InstanceID), http.StatusOK)
	displayResponse(body, &instance, "UpdateInstance")

	body = makeRequest(requests.GetInstance(instance.InstanceID), http.StatusOK)
	displayResponse(body, &instance, "GetInstance")

	body = makeRequest(requests.LinkToCollection(instance.Links.Version.HRef), http.StatusOK)
	displayResponse(body, &dataset, "LinkToCollection")

	body = makeRequest(requests.PublishDataset(instance.Links.Version.HRef), http.StatusOK)
	displayResponse(body, &dataset, "PublishDataset")

	log.Info("import dataset completed successfully", log.Data{
		"uri": cfg.PublishedDatasetURL(dataset.ID),
	})
}

func displayResponse(body []byte, model interface{}, requestName string) {
	json.Unmarshal(body, model)
	PrettyLogJSONResponse(body, model, requestName)
}

func makeRequest(req *http.Request, expectedStatus int) []byte {
	resp, err := client.Do(req)
	if err != nil {
		exit(err)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)

	if err := checkResponseStatus(expectedStatus, resp.StatusCode, b); err != nil {
		exit(err)
	}
	return b

}

func DropDatabases() {
	log.Info("dropping mongo databases", nil)
	sess, err := mgo.Dial(cfg.MongoURL)
	if err != nil {
		exit(err)
	}
	defer sess.Close()

	dbs, _ := sess.DatabaseNames()
	log.Debug("databases", log.Data{"": dbs})

	log.Debug("dropping imports", nil)
	if err := sess.DB("imports").DropDatabase(); err != nil {
		log.Error(err, nil)
	}

	log.Debug("dropping datasets", nil)
	if err := sess.DB("datasets").DropDatabase(); err != nil {
		log.Error(err, nil)
	}

	log.Info("dropping neo4j databases", log.Data{"config": cfg})
	pool, err := bolt.NewDriverPool(cfg.Neo4jURL, 1)
	if err != nil {
		exit(err)
	}
	conn, err := pool.OpenPool()
	if err != nil {
		exit(err)
	}
	defer conn.Close()
	res, err := conn.ExecNeo("MATCH(n) DETACH DELETE n", nil)
	log.Debug("results", log.Data{
		"delete results": res.Metadata()["stats"],
	})
	log.Info("dropping databases complete", nil)
}

func checkResponseStatus(expected int, actual int, body []byte) error {
	if actual != expected {
		err := fmt.Errorf("incorrect status code: %d", actual)
		log.Error(err, log.Data{
			"resp": string(body),
		})
		return err
	}
	return nil
}

func PrettyLogJSONResponse(b []byte, model interface{}, requestName string) {
	if logResp {
		json.Unmarshal(b, &model)
		pretty, _ := json.MarshalIndent(model, "", "  ")
		log.Debug("response", log.Data{
			requestName: string(pretty),
		})
	}
}

func exit(err error) {
	log.Error(err, nil)
	os.Exit(1)
}
