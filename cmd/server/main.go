package main

import (
	"log"

	"github.com/aaclee/mkn-api/pkg/http"
	"github.com/aaclee/mkn-api/pkg/logger"
)

const (
	configPath = "./config/config.development.json"
)

func main() {
	// Getting server configs
	c, err := http.GetServerConfigs(configPath)
	if err != nil {
		log.Fatalf("Error procesing config file at: %s, %s\n", configPath, err)
	}

	// Setting up logger
	logger := logger.CreateLogger()

	// Create server configs
	config := http.Config{
		Log:  logger,
		Port: c.Port,
	}

	// Create server
	server := http.CreateServer(config)

	// Listen and Serve
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server could not start on port: %d\n", config.Port)
	}
}
