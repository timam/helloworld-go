package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w, "home.page.gohtml")
}

func About(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w, "about.page.gohtml")
}
