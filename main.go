package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", ponger).Methods("GET")
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", r)
}

func ponger(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pong!")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "pong"}`))
}
