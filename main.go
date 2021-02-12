package main

import (
	"encoding/json"
	"net/http"
)

type SomeStruct struct {
	Color	string	`json:"color"`
	Message string 	`json:"message"`
}

func gimmeColor(w http.ResponseWriter, r *http.Request) {
	data := SomeStruct{
		Color: "#fc5c65", //FUSION RED
		Message: "kubernetes is awesome",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", gimmeColor)
	http.ListenAndServe(":8080", nil)
}
