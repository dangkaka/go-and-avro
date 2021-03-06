package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/linkedin/goavro"
	"github.com/dangkaka/go-and-avro-example/avro"
	"os"
	"bufio"
	"time"
	"strconv"
	"bytes"
)

type Activity struct {
	Id   string
	Type string
	Data string
}

const (
	JsonPath   = "examples/json.json"
	GoavroPath = "examples/goavro.avro"
	GoGenAvroPath = "examples/gogen-avro.avro"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", writeData).Methods("POST")
	r.HandleFunc("/json", getAllInJson).Methods("GET")
	r.HandleFunc("/goavro", getAllInGoAvro).Methods("GET")
	r.HandleFunc("/gogenavro", getAllInGoGenAvro).Methods("GET")

	r.HandleFunc("/generate", generateData).Methods("POST")


	fmt.Println("Listening on port 9999")
	log.Fatal(http.ListenAndServe(":9999", r))
}

func generateData(w http.ResponseWriter, r *http.Request) {
	var max int
	var err error
	if val, ok := r.URL.Query()["max"]; ok {
		max, err = strconv.Atoi(val[0])
		if err != nil {
		}
	} else {
		max = 100
	}
	now := time.Now()
	secs := now.Unix()
	longString := "This is a long long long long long long long long long long long long long data "
	for i := 0; i < max; i++ {
		value := (longString + strconv.Itoa(int(secs)) + " - index: " + strconv.Itoa(i))
		activity := Activity{
			"sampleId",
			"Create",
			value,
		}
		writeDataInJson(activity)
		writeDataInAvro(activity)
		writeDataInGoGenAvro(value)
	}
}

func getAllInJson(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(JsonPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var activities []Activity
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var activityOut Activity
		if err := json.Unmarshal([]byte(scanner.Text()), &activityOut); err != nil {
			panic(err)
		}
		activities = append(activities, activityOut)
	}

	json.NewEncoder(w).Encode(activities)
}

func getAllInGoAvro(w http.ResponseWriter, r *http.Request) {
	codec, err := goavro.NewCodec(`
        {
          "type": "record",
          "name": "Activity",
          "fields" : [
             {"name": "Id", "type": "string"},
			 {"name": "Type", "type": "string"},
			 {"name": "Data", "type": "string"}
          ]
        }`)
	f, err := os.Open(GoavroPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var activities []Activity
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		native, _, err := codec.NativeFromBinary([]byte(scanner.Text()))
		if err != nil {
			panic(err)
		}

		activityOut := StringMapToUser(native.(map[string]interface{}))

		activities = append(activities, activityOut)
	}

	json.NewEncoder(w).Encode(activities)
}

func getAllInGoGenAvro(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(GoGenAvroPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var activities []avro.Activity
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		activityOut, err := avro.DeserializeActivity(bytes.NewReader([]byte(scanner.Text())))
		if err != nil {
			panic(err)
		}

		activities = append(activities, *activityOut)
	}

	json.NewEncoder(w).Encode(activities)
}

func writeData(w http.ResponseWriter, r *http.Request) {
	var value string
	if val, ok := r.URL.Query()["value"]; ok {
		value = val[0]
	}
	activity := Activity{
		"1",
		"Create",
		value,
	}
	writeDataInJson(activity)
	writeDataInAvro(activity)
	writeDataInGoGenAvro(value)
}

func writeDataInJson(activity Activity) {
	json, err := json.Marshal(activity)
	if err != nil {
	}

	file, err := os.OpenFile(JsonPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660);
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, string(json))
	fmt.Fprintf(file, "\n")
}

func writeDataInAvro(activity Activity) {
	codec, err := goavro.NewCodec(`
        {
          "type": "record",
          "name": "Activity",
          "fields" : [
             {"name": "Id", "type": "string"},
			 {"name": "Type", "type": "string"},
			 {"name": "Data", "type": "string"}
          ]
        }`)
	if err != nil {
		fmt.Println(err)
	}
	// Convert native Go form to binary Avro data
	binary, err := codec.BinaryFromNative(nil, activity.ToStringMap())
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.OpenFile(GoavroPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660);
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, string(binary))
	fmt.Fprintf(file, "\n")
}

func writeDataInGoGenAvro(value string) {
	activity := avro.Activity{
		"1",
		"Create",
		value,
	}
	var buf bytes.Buffer
	activity.Serialize(&buf)
	file, err := os.OpenFile(GoGenAvroPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660);
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, buf.String())
	fmt.Fprintf(file, "\n")
}

func (activity *Activity) ToStringMap() map[string]interface{} {
	datumIn := map[string]interface{}{
		"Id": string(activity.Id),
		"Type":  string(activity.Type),
		"Data":  string(activity.Data),
	}
	return datumIn
}

func StringMapToUser(data map[string]interface{}) Activity {
	ind := Activity{}
	for k, v := range data {
		switch k {
		case "Id":
			if value, ok := v.(string); ok {
				ind.Id = value
			}
		case "Type":
			if value, ok := v.(string); ok {
				ind.Type = value
			}
		case "Data":
			if value, ok := v.(string); ok {
				ind.Data = value
			}
		}
	}
	return ind
}
