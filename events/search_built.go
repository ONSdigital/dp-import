package events

import "github.com/ONSdigital/dp-kafka/v2/avro"

// SearchIndexBuilt contains data related to a search that has just been built.
type SearchIndexBuilt struct {
	InstanceID    string `avro:"instance_id"`
	DimensionName string `avro:"dimension_name"`
}

var searchIndexBuilt = `{
  "type": "record",
  "name": "search-index-built",
  "fields": [
    {"name": "instance_id", "type": "string", "default": ""},
    {"name": "dimension_name", "type": "string", "default": ""}
  ]
}`

// SearchIndexBuiltSchema is the Avro schema for each dimension hierarchy successfuly sent to elastic
var SearchIndexBuiltSchema = &avro.Schema{
	Definition: searchIndexBuilt,
}
