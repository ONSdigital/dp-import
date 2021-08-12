package events

import "github.com/ONSdigital/dp-kafka/v2/avro"

// CantabularDatasetInstanceComplete is the event produced when the
// dimension options have been imported and the instance has successfully
// been imported from Cantabular
type CantabularDatasetInstanceComplete struct {
	InstanceID     string `avro:"instance_id"`
	CantabularBlob string `avro:"cantabular_blob"`
}

var cantabularDatasetInstanceCompleteSchema = `{
  "type": "record",
  "name": "cantabular-dataset-instance-complete",
  "fields": [
    {"name": "instance_id",     "type": "string"},
    {"name": "cantabular_blob", "type": "string"}
  ]
}`

// CantabularDatasetInstanceCompleteSchema is the avro schema for the
// CantabularDatasetInstanceComplete event
var CantabularDatasetInstanceCompleteSchema = &avro.Schema{
	Definition: cantabularDatasetInstanceCompleteSchema,
}
