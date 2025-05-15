package main

import (
	"html/template"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("hello").Parse(`<h1>Halo {{.}}</h1>`))
	tmpl.Execute(w, "Gopher")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
