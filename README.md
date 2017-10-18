dp-import
================

### Getting started

#### Required dependencies
 
##### Kafka

- Requires [Kafka 0.10.2.1](https://www.apache.org/dyn/closer.cgi?path=/kafka/0.10.2.1/kafka_2.11-0.10.2.1.tgz) 
- This version is not available via brew. Follow the [quickstart](https://kafka.apache.org/quickstart)

##### Neo4j

- Run ```brew install neo4j```
- Disable authentication by editing /usr/local/Cellar/neo4j/3.2.0/libexec/conf/neo4j.conf
  - Set ```dbms.security.auth_enabled=false```
- Run ```brew services restart neo4j```

##### MongoDB

* Run ```brew install mongodb```
* Run ```brew services start mongodb```

#### Running the services

There are a number of services that provide the import process functionality. Consider using [websysd](https://github.com/ONSdigital/dp/tree/master/websysd) 
that will run all of the services for you. If you do want to run services independently:

##### Import frontend
 - https://github.com/ONSdigital/dp-frontend-router `make debug`
 - https://github.com/ONSdigital/dp-frontend-dataset-controller `make debug`
 - https://github.com/ONSdigital/dp-frontend-filter-dataset-controller `make debug`
 - https://github.com/ONSdigital/zebedee `./run.sh` (run-reader on web)
 - https://github.com/ONSdigital/babbage `./run.sh`
 - https://github.com/ONSdigital/sixteens `./run.sh`
 - https://github.com/ONSdigital/florence `make debug`

##### Import backend

 - https://github.com/ONSdigital/dp-import-api
 - https://github.com/ONSdigital/dp-import-tracker
 - https://github.com/ONSdigital/dp-dimension-extractor
 - https://github.com/ONSdigital/dp-dimension-importer
 - https://github.com/ONSdigital/dp-observation-extractor
 - https://github.com/ONSdigital/dp-observation-importer
 - https://github.com/ONSdigital/dp-recipe-api

### Import a dataset

Navigate to `http://localhost:8081/florence/datasets/`
 - upload a file

Get instance data from the import API:
```curl localhost:21800/instances/284ca658-bfcf-4886-adc6-a43c4c040ad4 | jq .```

### Admin

#### Clear neo4j
```
brew services stop neo4j
rm -rf /usr/local/Cellar/neo4j/3.2.0/libexec/data
brew services start neo4j
```

#### clear kafka topic
   - enable kafka topic deletions: edit `/usr/local/etc/kafka/server.properties` and set `delete.topic.enable=true`

```kafka-topics --delete --zookeeper localhost:2181 --topic observation-extracted```

#### Create topic with partition / replication properties

```kafka-topics --create --zookeeper localhost:2181 --topic dimensions-inserted --partitions 1 --replication-factor 1```

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2016-2017, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
