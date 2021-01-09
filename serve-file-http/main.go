package main

import (
	"io"
	"net/http"
	"os"
)

func solo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="pisa.jpeg">`)
}

func pisa(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("pisa.jpeg")
	if err != nil {
		http.Error(w, "File Not Found", 404)
		return
	}
	defer file.Close()

	// using io.Copy to write file to http.ResponseWriter
	io.Copy(w, file)
}

func main() {
	http.HandleFunc("/", solo)
	http.HandleFunc("/pisa.jpeg", pisa)
	http.ListenAndServe(":8080", nil)
}
