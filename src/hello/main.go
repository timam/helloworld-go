package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	Logger "logger"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
	log.WithFields(log.Fields{
		"api": "/",
	}).Info("Hello World")
}

func main() {
	Logger.ConfigLogger("info.log")


	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
	log.WithFields(log.Fields{
		"host": "localhost",
		"port":   "8888",
	}).Info("Go server started successfully")
}
