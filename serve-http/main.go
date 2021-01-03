package main

import (
	"fmt"
	"net/http"
)

type fido int

func (f fido) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome, FIDO here! Whats yo poisen?")
}

func main() {
	var f fido
	http.ListenAndServe(":8080", f)
}
