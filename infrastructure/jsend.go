package infrastructure

import (
	"net/http"
	"encoding/json"
	"vendor/gorilla/schema"
)

// TO DO: Write tests using http://golang.org/pkg/net/http/httptest/

var decoder = schema.NewDecoder() // global decoder that caches struct reflections

type JSendResponse struct {
	Status	string
	Message	string
	Data	interface{}
}

func(j *JSendResponse) SendError(err error, w http.ResponseWriter) {
	j.Status = "error"
  j.Message = err.Error()
  b, _ := json.Marshal(j)
	w.Write(b)
}

func(j *JSendResponse) SendData(data interface{}, w http.ResponseWriter) {
	j.Status = "success"
  j.Data = data
  b, _ := json.Marshal(j)
	w.Write(b)
}