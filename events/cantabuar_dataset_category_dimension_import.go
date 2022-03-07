package events

import "github.com/ONSdigital/dp-kafka/v2/avro"

// CantabularDatasetCategoryDimensionImport is an event produced to trigger
// the import of dimension options for a given dimension from Cantabular
type CantabularDatasetCategoryDimensionImport struct {
	JobID          string `avro:"job_id"`
	InstanceID     string `avro:"instance_id"`
	DimensionID    string `avro:"dimension_id"`
	CantabularBlob string `avro:"cantabular_blob"`
	IsGeography    bool   `avro:"is_geography"`
}

var cantabularDatasetCategoryDimensionImportSchema = `{
  "type": "record",
  "name": "cantabular-dataset-category-dimension-import",
  "fields": [
    {"name": "dimension_id",   "type": "string"},
    {"name": "job_id", "type": "string"},
    {"name": "instance_id", "type": "string"},
    {"name": "cantabular_blob", "type": "string"}
    {"name": "is_geography", "type": "bool"}
  ]
}`

// CantabularDatasetCategoryDimensionImportSchema is the avro schema for the
// CantabularDatasetCategoryDimensionImport event
var CantabularDatasetCategoryDimensionImportSchema = &avro.Schema{
	Definition: cantabularDatasetCategoryDimensionImportSchema,
}
