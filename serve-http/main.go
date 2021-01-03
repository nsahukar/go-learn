package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
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

	data := struct {
		Host          string
		Method        string
		URL           *url.URL
		Header        http.Header
		ContentLength int64
		Submissions   map[string][]string
	}{
		r.Host,
		r.Method,
		r.URL,
		r.Header,
		r.ContentLength,
		r.Form,
	}
	tmpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	var f fido
	http.ListenAndServe(":8080", f)
}
