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

	fileStat, err := file.Stat()
	if err != nil {
		http.Error(w, "File Not Found", 404)
		return
	}

	// serving file with http.serveContent
	http.ServeContent(w, r, fileStat.Name(), fileStat.ModTime(), file)
}

func main() {
	http.HandleFunc("/", solo)
	http.HandleFunc("/pisa.jpeg", pisa)
	http.ListenAndServe(":8080", nil)
}
