package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"reflect"
	"sort"
	"time"
)

//Store the objects with map
var m  = make(map[string]object)

//Json Payload
type object struct {
	Timestamp string    `json:"timestamp"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`

}

//Convert map to slice
func mapToSlice(m map[string]object) []object {
	// Convert map to slice of keys.
	s := []object{}
	for _, value := range m {
		s = append(s, value)
	}
	//Sort it based on the timestamp
	sort.Slice(s, func(i, j int) bool { return s[i].Timestamp < s[j].Timestamp})

	return s
}

//Store the received objects by key
func addObject(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var obj object

	err := decoder.Decode(&obj)
	if err != nil {
		panic(err)
	}
	//Timestamp of the object
	obj.Timestamp = time.Now().Format("2006-01-02T15:04:05Z")
	//Store the object
	m[obj.Key] = obj

	w.Write([]byte("Success"))
}

//List all objects in descending order by timestamp
func listObjects(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		js, err := json.Marshal(mapToSlice(m))
		if err != nil {
			panic(err)
		}
		typeJs := reflect.TypeOf(js).Kind()
		log.Print(typeJs)
		w.Write(js)
}

func main() {
	//Set up a new router
	r := chi.NewRouter()
	//Initial Endpoint
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This server is developed by Owen for the BGL pre-interview task"))
	})
	//'/add' Endpoint
	r.Post("/add", addObject)

	//'/list' Endpoint
	r.Get("/list", listObjects)


	log.Fatal(http.ListenAndServe(":80", r))
}
