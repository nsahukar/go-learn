package main

import (
	"io"
	"net/http"
)

type fido int

func (f fido) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		io.WriteString(w, "Hello")
	case "/dog":
		io.WriteString(w, "Woof!")
	case "/cat":
		io.WriteString(w, "Meow...")
	}
}

func main() {
	var f fido
	http.ListenAndServe(":8080", f)
}
