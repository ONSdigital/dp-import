package events

import "github.com/ONSdigital/dp-kafka/v2/avro"

// CantabularDatasetInstanceStarted is an event produced when a cantabular import is triggered
type CantabularDatasetInstanceStarted struct {
	InstanceID     string `avro:"instance_id"`
	JobID          string `avro:"job_id"`
	CantabularType string `avro:"cantabular_type"`
}

var cantabularDatasetInstanceStartedSchema = `{
  "type": "record",
  "name": "cantabular-dataset-instance-started",
  "fields": [
    {"name": "instance_id", "type": "string"},
    {"name": "job_id", "type": "string"},
    {"name": "cantabular_type", "type": "string"}
  ]
}`

// CantabularDatasetInstanceStartedSchema provides an Avro schema for the CantabularDatasetInstanceStarted event
var CantabularDatasetInstanceStartedSchema = &avro.Schema{
	Definition: cantabularDatasetInstanceStartedSchema,
}
