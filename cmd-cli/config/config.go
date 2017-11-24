package config

import (
	"fmt"
	"github.com/ONSdigital/go-ns/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Config Model

type Model struct {
	DatasetAPIHost  string `yaml:"dataset-api-host"`
	ImportAPIHost   string `yaml:"import-api-host"`
	DevAuthToken    string `yaml:"dev-auth-token"`
	RecipeID        string `yaml:"recipe-id"`
	UploadAliasName string `yaml:"upload-alias-name"`
	UploadURL       string `yaml:"upload-url"`
	WebsiteHost     string `yaml:"website-host"`
	MongoURL        string `yaml:"mongo-url"`
	Neo4jURL        string `yaml:"neo4j-url"`
}

func Load() Model {
	source, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	var config Model
	if err := yaml.Unmarshal(source, &config); err != nil {
		panic(err)
	}

	Config = config
	log.Debug("cmd-cli configuration", log.Data{"": config})
	return config
}

func (c *Model) CreateDatasetURL(id string) string {
	return fmt.Sprintf("%s/datasets/%s", c.DatasetAPIHost, id)
}

func (c *Model) GetDatasetURL(id string) string {
	return fmt.Sprintf("%s/datasets/%s", c.DatasetAPIHost, id)
}

func (c *Model) GetInstancesURL() string {
	return fmt.Sprintf("%s/instances", c.DatasetAPIHost)
}

func (c *Model) GetInstanceURL(id string) string {
	return fmt.Sprintf("%s/instances/%s", c.DatasetAPIHost, id)
}

func (c *Model) UpdateInstanceURL(id string) string {
	return fmt.Sprintf("%s/instances/%s", c.DatasetAPIHost, id)
}

func (c *Model) UploadFileURL() string {
	return fmt.Sprintf("%s/jobs", c.ImportAPIHost)
}

func (c *Model) SubmitJob(id string) string {
	return fmt.Sprintf("%s/jobs/%s", c.ImportAPIHost, id)
}

func (c *Model) PublishedDatasetURL(id string) string {
	return fmt.Sprintf("%s/datasets/%s", c.WebsiteHost, id)
}
