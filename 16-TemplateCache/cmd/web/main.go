package main

import (
	"fmt"
	"github.com/timam/go-mod/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"



func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
