#### Scripts

### dp-import

Access APIs, mongo, neo in a nicely-wrapped script:

- API access (currently ImportAPI and DatasetAPI)
- Query, update, destroy your databases (currently MongoDB, Neo4J)
- send kafka messages

Typical usage:

```sh
$ PATH=$PATH:$GOPATH/src/github.com/ONSdigital/dp-import/scripts
$ dp-import db_wipe
# ...delete all data from your databases...
$ dp-import create
# ...will ask you to confirm POSTing JSON to the ImportAPI...
# ...shows you the resultant import-job and dataset-instance...
$ dp-import db
# ...query the database for import-jobs, instances, even neo...
$ dp-import jobs
# ...get the list of import-jobs from the ImportAPI...
$ dp-import inst
# ...get the list of instances from the DatasetAPI...
$ dp-import inst xyz
# ...get the instance whose ID starts 'xyz' from the DatasetAPI...

$ dp-import submit abc
# ...change the import-job whose ID starts 'abc' to have state: submitted...
```
