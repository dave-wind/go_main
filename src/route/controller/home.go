package controller

import (
	"net/http"
	"text/template"
)

func registerHomeRoutes() {
	http.HandleFunc("/home", HanleHome)
}

func HanleHome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "home.html")
	t.ExecuteTemplate(w, "layout", "Hello World")
}
