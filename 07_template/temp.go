package main

import (
	"html/template"
	"net/http"
)

func temp(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	t, _ := template.ParseFiles("temp.html")
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", temp)
	http.ListenAndServe(":8080", nil)
}
