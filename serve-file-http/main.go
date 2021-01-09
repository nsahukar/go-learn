package main

import (
	"io"
	"net/http"
)

func solo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="pisa.jpeg">`)
}

func pisa(w http.ResponseWriter, r *http.Request) {
	// serving file with http.serveFile
	http.ServeFile(w, r, "pisa.jpeg")
}

func main() {
	http.HandleFunc("/", solo)
	http.HandleFunc("/pisa.jpeg", pisa)
	http.ListenAndServe(":8080", nil)
}
