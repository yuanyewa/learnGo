package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from Go!")
}

func world(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "oh my God!")
}
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	http.ListenAndServe(":80", nil)
}
