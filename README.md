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


Ensure a dataset exists on the dataset API for the dataset specified in the recipe being used.
The current stubbed recipe api specifies dataset ID 931a8a2a-0dc8-42b6-a884-7b6054ed3b68 for the CPI dataset recipe.

```
curl --header 'internal-token:FD0108EA-825D-411C-9B1D-41EF7727F465' http://localhost:22000/datasets/931a8a2a-0dc8-42b6-a884-7b6054ed3b68
```

If not then create one (already set to published):
```
curl -X POST -d '{"release_frequency":"yearly", "state": "published", "theme": "population", "title": "CPI" }' --header 'internal-token:FD0108EA-825D-411C-9B1D-41EF7727F465' http://localhost:22000/datasets/931a8a2a-0dc8-42b6-a884-7b6054ed3b68
```

#### Create instance

Navigate to <http://localhost:8081/florence/datasets>
 - upload a file
 - select an edition
 - click submit to publishing

Get instance data from the import API - the instance state should be 'complete' if the import succeeded (copy the instance ID - it will be the last instance in the array):

Example curl command to GET instances:
```
curl --header 'internal-token:FD0108EA-825D-411C-9B1D-41EF7727F465' http://localhost:22000/instances | jq
```

#### Set the release date and license values on the instance (replace the instance ID with the one you created).

API call details:
```
PUT localhost:22000/instances/750102f4-2839-441f-b2e4-6cf99d26858a
{
	"release_date": "todayisfine",
	"license": "todayisfine"
}
```

Example curl command to PUT instance data:
```
curl -v -X PUT -d '{"release_date":"today", "license": "wut"}' --header 'internal-token:FD0108EA-825D-411C-9B1D-41EF7727F465' localhost:22000/instances/750102f4-2839-441f-b2e4-6cf99d26858a
```

#### Set the instance to 'edition-confirmed' (replace the instance ID with the one you created)

This converts from an instance to a dataset version.

API call details:
```
PUT localhost:22000/instances/750102f4-2839-441f-b2e4-6cf99d26858a
{
    "edition":"Time-series",
    "state": "edition-confirmed"
}
```

Example curl command to PUT instance state:
```
curl -v -X PUT -d '{"state":"edition-confirmed", "edition":"Time-series"}' --header 'internal-token:FD0108EA-825D-411C-9B1D-41EF7727F465' localhost:22000/instances/750102f4-2839-441f-b2e4-6cf99d26858a
```

#### Associate the dataset with a collection

You will first need to get the dataset version URL from the instances endpoint

Example curl command to GET instances:
```
curl --header 'internal-token:FD0108EA-825D-411C-9B1D-41EF7727F465' http://localhost:22000/instances | jq
```

Copy the link for the version URL to use in the following calls.

API call details:
```
PUT http://localhost:22000/datasets/931a8a2a-0dc8-42b6-a884-7b6054ed3b68/editions/Time-series/versions/1
{
	"collection_id": "1234",
	"state": "associated"
}
```

Example curl command to PUT dataset version state / collection:
```
curl -v -X PUT -d '{"state":"associated", "collection_id":"123"}' --header 'internal-token:FD0108EA-825D-411C-9B1D-41EF7727F465' http://localhost:22000/datasets/931a8a2a-0dc8-42b6-a884-7b6054ed3b68/editions/Time-series/versions/1
```

#### Set dataset to published

API call details:
```
PUT http://localhost:22000/datasets/931a8a2a-0dc8-42b6-a884-7b6054ed3b68/editions/Time-series/versions/1
{
	"state": "published"
}
```

Example curl command to PUT dataset version state:
```
curl -v -X PUT -d '{"state":"published"}' --header 'internal-token:FD0108EA-825D-411C-9B1D-41EF7727F465' http://localhost:22000/datasets/931a8a2a-0dc8-42b6-a884-7b6054ed3b68/editions/Time-series/versions/1
```

#### Check dataset is available via the frontend

You should now see the details of the dataset via the frontend by navigating to <http://localhost:20000/datasets/931a8a2a-0dc8-42b6-a884-7b6054ed3b68>

For filtering hierarchical dimensions, ensure you have created the hierarchy for the instance as detailed here: <https://github.com/ONSdigital/dp-hierarchy-builder>

### Admin

#### Clear MongoDB

The following will drop the datasets and imports databases. Add further databases as required.

```
 mongo mongodb://localhost:27017 <<EOF
 use datasets
 db.dropDatabase();
 use imports
 db.dropDatabase();
EOF
```

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
