package main

import (
	"fmt"
	"github.com/timam/go-mod/pkg/config"
	"github.com/timam/go-mod/pkg/handlers"
	"github.com/timam/go-mod/pkg/render"
	"net/http"
)

const portNumber = ":8080"



func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil{
		fmt.Println("Can't create template cache")
	}
	app.TemplateCache = tc
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
