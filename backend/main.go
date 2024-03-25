package main

import (
	"log"
	"net/http"
)

func main() {
	// Set up a simple HTTP route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to LogSphere!"))
	})

	// Configure and start the HTTP server
	log.Println("Starting LogSphere server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
