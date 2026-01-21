package main

// api base

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func api_base() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
