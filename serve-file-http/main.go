package main

import (
	"io"
	"net/http"
)

func solo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/The_Leaning_Tower_of_Pisa_SB.jpeg/1200px-The_Leaning_Tower_of_Pisa_SB.jpeg">`)
}

func main() {
	http.HandleFunc("/", solo)
	http.ListenAndServe(":8080", nil)
}
