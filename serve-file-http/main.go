package main

import (
	"io"
	"net/http"
)

func fido(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, User!")
}

func pisa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/img/pisa.jpeg">`)
}

func main() {
	http.HandleFunc("/", fido)
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("img"))))
	http.HandleFunc("/pisa", pisa)
	http.ListenAndServe(":8080", nil)
}
