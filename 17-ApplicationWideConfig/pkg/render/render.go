package render

import (
	"bytes"
	"fmt"
	"github.com/timam/go-mod/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)


var functions = template.FuncMap{}
var app *config.AppConfig

//NewTemplates sets the config for render package
func NewTemplates(a *config.AppConfig)  {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc := app.TemplateCache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template form TemplateCache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w)
	if err != nil{
		fmt.Println("Error writing template to browser", err)
	}
}

//CreateTemplateCache creates a template as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil{
		return myCache, err
	}
	for _, page := range pages{
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil{
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil{
			return myCache, err
		}
		if len(matches) > 0{
			ts, err := ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil{
				return myCache, err
			}
			myCache[name] = ts
		}

	}

	return myCache, nil
}
