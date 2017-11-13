# go-and-avro
Test go and avro, compare avro with json

h4. Generate sample data
Ex: generate 10000 rows
```curl -X POST 'http://localhost:9999/generate?max=10000'```

Generated data stays in 
```
const (
	JsonPath   = "/tmp/go-and-avro.json"
	AvroPath = "/tmp/go-and-avro.avro"
)
```

h4. View
* Avro
```http://localhost:9999/avro```

* Json
```http://localhost:9999/json```
