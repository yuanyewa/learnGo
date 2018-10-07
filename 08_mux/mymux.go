package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func handleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// w.Write([]byte(vars["user"]))
	fmt.Fprintf(w, vars["user"])
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/profile/{user}", handleUser)
	// http.Handle("/", r)
	http.Handle("/", handlers.LoggingHandler(os.Stdout, r))
	http.ListenAndServe(":8080", nil)
}
