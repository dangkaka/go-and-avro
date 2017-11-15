# go-and-avro
Test go and avro, compare avro with json

## What can we do here
- Be able to generate simple data (in json and avro) and save it in 2 files
- Be able to get data from json and avro from those 2 files by GET request

## Generate sample data
Ex: generate 10000 rows
```curl -X POST 'http://localhost:9999/generate?max=10000'```

Generated data stays in 
```
const (
	JsonPath   = "/tmp/go-and-avro.json"
	AvroPath = "/tmp/go-and-avro.avro"
)
```

## View
* Avro
```http://localhost:9999/avro```

* Json
```http://localhost:9999/json```
