package main

import (
	"encoding/json"
	"github.com/timam/helloworld-go/logger"
	"net/http"
)

type SomeStruct struct {
	Color	string	`json:"color"`
	Message string 	`json:"message"`
}

func colorAndMessage(w http.ResponseWriter, r *http.Request) {
	data := SomeStruct{
		Color: "#fc5c65", //FUSION RED
		Message: "kubernetes is awesome",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
	logger.Log.Printf("SUCCESS")
}

func main() {
	http.HandleFunc("/", colorAndMessage)
	http.ListenAndServe(":8080", nil)
}
