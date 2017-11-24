package requests

import (
	"bytes"
	"encoding/json"

	"github.com/ONSdigital/dp-dataset-api/models"
	imodel "github.com/ONSdigital/dp-import-api/models"
	"github.com/ONSdigital/dp-import/cmd-cli/config"
	"github.com/ONSdigital/go-ns/log"
	"io"
	"net/http"
)

var (
	Cfg config.Model

	LogRequests bool = false

	createDatasetJson, _ = json.Marshal(models.Dataset{
		ReleaseFrequency: "yearly",
		State:            "published",
		Theme:            "population",
		Title:            "CPI",
	})

	uploadFileJob, _ = json.Marshal(imodel.Job{
		RecipeID: "b944be78-f56d-409b-9ebd-ab2b77ffe187",
		UploadedFiles: &[]imodel.UploadedFile{imodel.UploadedFile{
			AliasName: "CPI COICOP",
			URL:       "s3://dp-dimension-extractor/EXAMPLE_V4-coicopcomb-inc-geo-code.csv",
		}},
	})

	submitJob, _ = json.Marshal(imodel.Job{
		State: "submitted",
	})

	updatedInstance, _ = json.Marshal(models.Instance{
		ReleaseDate: "todayisfine",
		Edition:     "Time-series",
		State:       "edition-confirmed",
	})

	versionAssociated, _ = json.Marshal(models.Version{
		CollectionID: "1234567890",
		State:        "associated",
	})

	publishVersion, _ = json.Marshal(models.Version{
		State: "published",
	})
)

func CreateDataset(datasetID string) *http.Request {
	return create("CreateDataset", http.MethodPost, Cfg.CreateDatasetURL(datasetID), createDatasetJson)
}

func GetDataset(datasetID string) *http.Request {
	return create("GetDataset", http.MethodGet, Cfg.GetDatasetURL(datasetID), nil)
}

func UploadFile() *http.Request {
	return create("UploadFile", http.MethodPost, Cfg.UploadFileURL(), uploadFileJob)
}

func Submit(jobID string) *http.Request {
	return create("SubmitJob", http.MethodPut, Cfg.SubmitJob(jobID), submitJob)
}

func GetInstances() *http.Request {
	return create("GetInstances", http.MethodGet, Cfg.GetInstancesURL(), nil)
}

func GetInstance(instanceID string) *http.Request {
	return create("GetInstance", http.MethodGet, Cfg.GetInstanceURL(instanceID), nil)
}

func UpdateInstance(instanceID string) *http.Request {
	return create("UpdateInstance", http.MethodPut, Cfg.UpdateInstanceURL(instanceID), updatedInstance)
}

func LinkToCollection(url string) *http.Request {
	return create("LinkToCollection", http.MethodPut, url, versionAssociated)
}

func PublishDataset(url string) *http.Request {
	return create("UpdateInstance", http.MethodPut, url, publishVersion)
}

func create(context string, method string, url string, body []byte) *http.Request {
	var reader io.Reader = nil
	if body != nil {
		reader = bytes.NewBuffer(body)
	}
	r, _ := http.NewRequest(method, url, reader)
	r.Header.Set("Internal-Token", Cfg.DevAuthToken)
	if LogRequests {
		log.Info("", log.Data{
			"request": context,
			"url":     r.URL.String(),
			"method":  r.Method,
		})
	}
	return r
}
