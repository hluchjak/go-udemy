package handlers

import (
	"net/http"
	"website/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.gohtml")
}

func About(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "This is the about page.")
	render.Template(w, "about.page.gohtml")
}
