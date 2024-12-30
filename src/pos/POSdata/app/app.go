package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler for the root route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to the POS System Backend!")
}

// Main function to set up the server
func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/", homeHandler).Methods("GET")

	// Start the server
	port := "8080"
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
