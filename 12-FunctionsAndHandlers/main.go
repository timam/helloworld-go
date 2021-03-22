package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request)  {
	_, _ = fmt.Fprintf(w, "this is the home page")
}
func About(w http.ResponseWriter, r *http.Request)  {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))
}

func addValues(x, y int) int {
	return x+y
}

func main()  {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Application started at port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}