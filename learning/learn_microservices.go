/*
	What are microservices?

Microservices are an architectural style that structures an application as a
collection of small, loosely coupled services. Each service is self-contained
and can be developed, deployed, and scaled independently. These services
communicate with each other through lightweight protocols such as HTTP or
message queues.


	Why use Golang for building microservices?

Golang is a great choice for building microservices due to its simplicity,
performance, and built-in support for concurrency. Its compiled nature makes
it efficient and allows for easy deployment. Golang's standard library
provides essential packages for building HTTP servers, handling JSON,
and working with databases.


	Creating a simple microservice with Golang

To get started, let's create a simple microservice that exposes an API
endpoint to retrieve user information. We will use the Gorilla Mux package
for routing and handling HTTP requests.
*/

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Simulate fetching user information from a database
	user := User{
		ID:    "1",
		Name:  "John Doe",
		Email: "john@example.com",
	}

	// Convert the user struct to JSON
	json.NewEncoder(w).Encode(user)
}

func main() {
	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define the API route
	router.HandleFunc("/user", GetUser).Methods("GET")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8000", router))
}
