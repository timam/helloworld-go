package handlers

import (
	"github.com/timam/go-mod/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	render.RenderTemplate(w, "home.page.gohtml")
}

func About(w http.ResponseWriter, r *http.Request)  {
	render.RenderTemplate(w, "about.page.gohtml")
}
