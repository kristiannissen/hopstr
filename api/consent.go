package api

import (
	"github.com/google/uuid"
	"log"
	"net/http"
)

type Message struct {
	status int    `json:statuscode`
	id     string `json:id`
}

func Consent(w http.ResponseWriter, req *http.Request) {
	//
	// w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")

	// Check request method
	if (req.Method == "GET") {
		// Return new uuid
		w.Header().Set("Content-type", "application/javascript")
		w.WriteHeader(http.StatusOK)
		u := uuid.New().String()
		w.Write([]byte("{uuid:\""+u+"\"};"))
	} else if (req.Method == "POST") {
		log.Println("Store new consent")
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"msg\": \"OK\"}"))
	} else {
		log.Println("Method not allowed")
	}
}
