package events

import "github.com/ONSdigital/dp-kafka/v3/avro"

// CantabularCSVWCreated event
type CantabularCSVWCreated struct {
	DatasetID  string `avro:"dataset_id"`
	Edition    string `avro:"edition"`
	Version    string `avro:"version"`
	InstanceID string `avro:"instance_id"`
	RowCount   int32  `avro:"row_count"`
}

var cantabularCSVWCreatedSchema = `{
  "type": "record",
  "name": "cantabular-metadata-complete",
  "fields": [
    {"name": "instance_id", "type": "string", "default": ""},
    {"name": "dataset_id",  "type": "string", "default": ""},
    {"name": "edition",     "type": "string", "default": ""},
    {"name": "version",     "type": "string", "default": ""},
    {"name": "row_count",   "type": "int", "default": 0}
  ]
}`

// CantabularCSVWCreatedSchema avro schema
var CantabularCSVWCreatedSchema = &avro.Schema{
	Definition: cantabularCSVWCreatedSchema,
}
