package main

import (
	"fmt"
	"net/http"
)

func myserv(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello ")
}

func main() {
	http.HandleFunc("/", myserv)
	http.ListenAndServe(":8080", nil)
}
