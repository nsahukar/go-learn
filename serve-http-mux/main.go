package main

import (
	"io"
	"net/http"
)

func fido(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Woof!")
}

func cat(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Meow...")
}

func main() {
	http.HandleFunc("/", fido)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)

	http.ListenAndServe(":8080", nil)
}
