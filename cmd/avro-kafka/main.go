package main

import (
	"flag"
	"os"
	"time"

	"github.com/ONSdigital/go-ns/avro"
	"github.com/ONSdigital/go-ns/kafka"
	"github.com/ONSdigital/go-ns/log"
)

// InputFileAvailableSchema schema
var InputFileAvailableSchema = `{
	"type": "record",
	"name": "input-file-available",
	"fields": [
		{"name": "file_url", "type": "string"},
		{"name": "instance_id", "type": "string"}
	]
}`

type inputFileAvailable struct {
	fileURL    string `avro:"file_url"`
	InstanceID string `avro:"instance_id"`
}

// from dp-observation-importer
var observationsInsertedSchema = `{
  "type": "record",
  "name": "import-observations-inserted",
  "fields": [
    {"name": "instance_id", "type": "string"},
    {"name": "observations_inserted", "type": "int"}
  ]
}`

type insertedObservationsMessage struct {
	InstanceID           string `avro:"instance_id"`
	ObservationsInserted int32  `avro:"observations_inserted"`
}

func main() {
	instance_id := flag.String("id", "21", "instance id")
	producerTopic := flag.String("topic", "input-file-available", "producer topic")
	fileURL := flag.String("s3", "s3://dp-dimension-extractor/OCIGrowth.csv", "s3 file")
	broker := flag.String("kafka", "localhost:9092", "kafka address")
	insertedObservations := flag.Int("inserts", 2000, "inserted observations")
	flag.Parse()

	producer, err := kafka.NewProducer([]string{*broker}, *producerTopic, int(2000000))
	if err != nil {
		panic(err)
	}

	var schema *avro.Schema
	var producerMessage []byte

	if *producerTopic == "input-file-available" {
		schema = &avro.Schema{Definition: InputFileAvailableSchema}
		producerMessage, err = schema.Marshal(&inputFileAvailable{
			//fileURL:    "s3://dp-dimension-extractor/UKBAA01a.csv",
			fileURL:    *fileURL,
			InstanceID: *instance_id,
		})
	} else if *producerTopic == "import-observations-inserted" {
		schema = &avro.Schema{Definition: observationsInsertedSchema}
		insertedObservationsMsg := insertedObservationsMessage{
			InstanceID:           *instance_id,
			ObservationsInserted: int32(*insertedObservations),
		}
		log.Debug("msg", log.Data{"iom": insertedObservationsMsg})
		producerMessage, err = schema.Marshal(&insertedObservationsMsg)
	}

	if err != nil {
		log.ErrorC("error marshalling", err, nil)
		panic(err)
		os.Exit(1)
	}

	producer.Output() <- producerMessage
	time.Sleep(time.Duration(1000 * time.Millisecond))
	producer.Closer() <- true
}
