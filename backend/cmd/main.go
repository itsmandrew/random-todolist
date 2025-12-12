package main

import (
	"log"
	"net/http"
	"time"

	"github.com/itsmandrew/random-todolist/internal/handlers"
)

func main() {
	// Create a new ServeMux (router)
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)

	// Define server configuration
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Starting the server
	log.Printf("Server starting on %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
	}
}
