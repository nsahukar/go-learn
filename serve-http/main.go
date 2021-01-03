package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type fido int

func init() {
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
}

func (f fido) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tmpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

func main() {
	var f fido
	http.ListenAndServe(":8080", f)
}
