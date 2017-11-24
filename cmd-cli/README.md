# CMD-CLI

dev tool for dropping both Mongo & Neo4j databases and freshly importing & publishing the CPI COICOP dataset.

### build
```bash
go build -o cmd-cli
``` 

### drop all databases
```bash
./cmd-cli -cmd=clean
```

### drop all databases & reimport/publish
```bash
./cmd-cli -cmd=clean-import
```
