package controller

import (
	"net/http"
	"text/template"
)

func registerAboutRoutes() {
	http.HandleFunc("/about", HanleAbout)
}

func HanleAbout(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "about.html")
	t.ExecuteTemplate(w, "layout", "Hello World")
}
