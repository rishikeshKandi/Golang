package handler

import (
	"net/http"

	"github.com/rishikeshKandi/Golang.git/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.html")
}
