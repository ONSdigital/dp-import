# *cmd-cli*

dev tool for dropping both Mongo and Neo4j databases and freshly importing and publishing the CPI COICOP dataset
and generating the necessary hierarchy data.


### Getting started
git clone to you `GOPATH`
```bash
git clone git@github.com:ONSdigital/dp-import.git
```

### Dependencies
**cmd-cli** uses [dp-hierarchy-builder](https://github.com/ONSdigital/dp-hierarchy-builder) to generate hierarchies in 
an unconventional ~~hacky~~ MVP sort of way so it needs to be available on your `GOPATH`

```bash
git clone git@github.com:ONSdigital/dp-hierarchy-builder.git
```

### build
```bash
go build -o cmd-cli
``` 

### clean
To drop all MongoDB databases and empty Neo4j
```bash
./cmd-cli -cmd=clean
```

### clean-import
To drop all MongoDB databases, empty Neo4j, freshly import and publish a dataset and generate the required hierarchy data
```bash
./cmd-cli -cmd=clean-import
```
