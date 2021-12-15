package events

import "github.com/ONSdigital/dp-kafka/v3/avro"

// CantabularCSVCreated event
type CantabularCSVCreated struct {
	InstanceID string `avro:"instance_id"`
	DatasetID  string `avro:"dataset_id"`
	Edition    string `avro:"edition"`
	Version    string `avro:"version"`
	RowCount   int32  `avro:"row_count"`
}

var cantabularCSVCreatedSchema = `{
  "type": "record",
  "name": "cantabular-csv-created",
  "fields": [
    {"name": "instance_id", "type": "string", "default": ""},
    {"name": "dataset_id", "type": "string", "default": ""},
    {"name": "edition", "type": "string", "default": ""},
    {"name": "version", "type": "string", "default": ""},
    {"name": "row_count", "type": "int", "default": 0}
  ]
}`

// CantabularCSVCreatedSchema avro schema
var CantabularCSVCreatedSchema = &avro.Schema{
	Definition: cantabularCSVCreatedSchema,
}
