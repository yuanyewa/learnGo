package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func mymiddle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("this is me")
		r = r.WithContext(context.WithValue(r.Context(), "foo", "bar"))
		next(w, r)
	}
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
	// http.Handle("/", r)
	http.Handle("/", handlers.LoggingHandler(os.Stdout, r))
	http.ListenAndServe(":8080", nil)
}
