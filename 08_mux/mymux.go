package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type myType struct {
	Name string `json:"name"`
	Age  int    `json:"ageis"`
}

func mymiddle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("this is me")
		r = r.WithContext(context.WithValue(r.Context(), "foo", "bar"))
		next(w, r)
	}
}

func restWithName(w http.ResponseWriter, r *http.Request) {
	// name := r.URL.Query().Get("name")
	log.Println("I'm insie here")
	v := mux.Vars(r)
	t := myType{Name: v["name"], Age: 16}
	json.NewEncoder(w).Encode(t)
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// w.Write([]byte(vars["user"]))
	fmt.Fprintf(w, vars["user"])
	fmt.Fprintf(w, r.Context().Value("foo").(string))
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/profile/{user}", mymiddle(handleUser))
	r.HandleFunc("/rest/{name}", restWithName).Methods("GET")
	// http.Handle("/", r)
	http.Handle("/", handlers.LoggingHandler(os.Stdout, r))
	http.ListenAndServe(":8080", nil)
}
