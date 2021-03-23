package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w, "home.page.gohtml")
}

func About(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w, "about.page.gohtml")
}

func renderTemplate(w http.ResponseWriter, tmpl string)  {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("| Error | Parsing Template |", err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
