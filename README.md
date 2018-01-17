# go-and-avro-example
Test go and avro, compare json and 2 avro libraries
https://github.com/alanctgardner/gogen-avro
https://github.com/linkedin/goavro

## What can we do here
- Be able to generate simple data (in json and 2 avro libraries) and save them in 3 files
- Be able to get data from those 3 files by GET request

## Generate sample data
Ex: generate 10000 rows
```curl -X POST 'http://localhost:9999/generate?max=10000'```

Generated data stays in 
```
const (
	JsonPath   = "examples/json.json"
	GoavroPath = "examples/goavro.avro"
	GoGenAvroPath = "examples/gogen-avro.avro"
)
```

## View
* Json
```http://localhost:9999/json```

* Goavro with 
```http://localhost:9999/goavro```

* GoGenAvro
```http://localhost:9999/gogenavro```

