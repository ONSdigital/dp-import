package events

import "github.com/ONSdigital/dp-kafka/v3/avro"

// CantabularExportStart event
type CantabularExportStart struct {
	InstanceID string `avro:"instance_id"`
	DatasetID  string `avro:"dataset_id"`
	Edition    string `avro:"edition"`
	Version    string `avro:"version"`
}

var cantabularExportStartSchema = `{
  "type": "record",
  "name": "cantabular-export-start",
  "fields": [
    {"name": "instance_id", "type": "string", "default": ""},
    {"name": "dataset_id",  "type": "string", "default": ""},
    {"name": "edition",     "type": "string", "default": ""},
    {"name": "version",     "type": "string", "default": ""}
  ]
}`

// CantabularExportStartSchema avro schema
var CantabularExportStartSchema = &avro.Schema{
	Definition: cantabularExportStartSchema,
}
