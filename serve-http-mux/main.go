package main

import (
	"io"
	"net/http"
)

type fido int
type hotdog int
type hotcat int

func (f fido) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello")
}

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Woof!")
}

func (c hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Meow...")
}

func main() {
	var f fido
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/", f)
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
