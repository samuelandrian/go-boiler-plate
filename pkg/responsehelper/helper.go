package responsehelper

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	if data == nil {
		data = map[string]interface{}{}
	}
	resp, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp)
}

func SendErrorResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	HTTPResponse(w, data, statusCode)
}

func HTTPResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	responseTemplate := ResponseStruct{
		Data: data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	resp, err := json.Marshal(responseTemplate)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp)
}

type ResponseStruct struct {
	Data interface{} `json:"result"`
}
