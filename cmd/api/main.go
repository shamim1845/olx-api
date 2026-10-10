package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// create a new servemux
	mux := http.NewServeMux()

	// Root API endpoint
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{"version": "1.0.0", "status": "ok", "message": "Welcome to OLX clone API"}`))
	})

	// Health check endpoint
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{"status": "ok", "message": "Server is up and running."}`))
	})

	// Get port from environment variable or default to 8090
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8090"
	}
	
	// create HTTP server with timeout
	srv := http.Server {
		Addr: ":" + PORT,
		Handler: mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	
	// Log before blocking call
	log.Printf("Server is running on port %v", PORT)

	// Listen on a specific port (blocks until error)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server can't start: %v", err)
	}
}
