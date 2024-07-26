package handlers

import (
	"main/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to the home page")
	render.RenderTemplate(w, "home.page.tmpl")

}

func About(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to the About page and ")
	render.RenderTemplate(w, "about.page.tmpl")
}
