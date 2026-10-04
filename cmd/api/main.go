package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	// Root API endpoint
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Welcome to OLX clone API"}`))
	})

	// Health check endpoint
	http.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "Server is up and running.",
		})
	})

	// Get port from environment variable or default to 8090
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8090"
	}

	// Log before blocking call
	log.Printf("Server is running on port %v", PORT)

	// Listen on a specific port (blocks until error)
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatalf("Server can't start: %v", err)
	}
}
