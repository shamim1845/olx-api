package main

import (
	"log"
	"net/http"
	"time"

	"github.com/shamim1845/olx-api/internal/config"
)

func main() {
	// load configuration
	cfg := config.MustLoad()

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

	// create HTTP server with timeout
	srv := http.Server {
		Addr: ":" + cfg.Port,
		Handler: mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	
	// Log before blocking call
	log.Printf("Server is running on port %v", cfg.Port)

	// Listen on a specific port (blocks until error)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server can't start: %v", err)
	}
}
