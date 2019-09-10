package main

import (
	"log"

	"github.com/aaclee/mkn-api/pkg/http"
	"github.com/aaclee/mkn-api/pkg/logger"
)

func main() {
	// Setting up logger
	logger := logger.CreateLogger()

	// Create server configs
	config := http.Config{
		Log:  logger,
		Port: 8000,
	}

	// Create server
	server := http.CreateServer(config)

	// Listen and Serve
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server could not start on port: %d\n", config.Port)
	}
}
