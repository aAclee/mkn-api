package main

import (
	"encoding/json"
	"flag"
	"log"
	stdHttp "net/http"

	"github.com/aaclee/mkn-api/pkg/auth"
	"github.com/aaclee/mkn-api/pkg/http"
	"github.com/aaclee/mkn-api/pkg/logger"
	"github.com/aaclee/mkn-api/pkg/player"
	"github.com/aaclee/mkn-api/pkg/postgres"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

const (
	serverConfig   = "./config/server.development.json"
	postgresConfig = "./config/postgres.development.json"
)

func main() {
	// Getting server configs
	c, err := http.GetServerConfigs(serverConfig)
	if err != nil {
		log.Fatalf("Error processing server config file at: %s, %s\n", serverConfig, err)
	}

	// Setting up logger
	logger := logger.CreateLogger()

	// Getting database configs
	dc, err := postgres.GetDatabaseConfigs(postgresConfig)
	if err != nil {
		log.Fatalf("Error processing database config file at: %s, %s\n", serverConfig, err)
	}

	// Override flags
	var dbHost = flag.String("dbhost", dc.Host, "override database hostname")

	flag.Parse()

	// Create postgres connection
	db, err := postgres.CreateConnection(postgres.Config{
		Host:     *dbHost,
		Port:     dc.Port,
		User:     dc.User,
		Password: dc.Password,
		DBname:   dc.Database,
	})
	if err != nil {
		log.Fatalf("Error creating postgres connection: %s", err)
	}

	// Create server configs
	config := http.Config{
		Log:  logger,
		Port: c.Port,
	}

	// Create server
	server := http.CreateServer(config)

	// Create server mux
	r := mux.NewRouter()

	// Mux middleware
	r.Use(middleware)

	// Repositories
	authRepository := auth.CreatePostgresRepository(db)
	playerRepository := player.CreatePostgresRepository(db)

	// Services
	authService := auth.CreateService(authRepository, playerRepository)

	// Handlers
	authHandler := auth.CreateHandler(authService)
	r.HandleFunc("/api/auth", authHandler.Authenticate).Methods("POST")

	// Handler for non-existing routes
	r.PathPrefix("/").HandlerFunc(catchAllHandler)

	// Listen and Serve
	err = server.ListenAndServe(r)
	if err != nil {
		log.Fatalf("Server could not start on port: %d\n", config.Port)
	}
}

func catchAllHandler(w stdHttp.ResponseWriter, r *stdHttp.Request) {
	w.WriteHeader(stdHttp.StatusNotFound)

	json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{
		Error: "Path not found!",
	})
}

func middleware(next stdHttp.Handler) stdHttp.Handler {
	return stdHttp.HandlerFunc(func(w stdHttp.ResponseWriter, r *stdHttp.Request) {
		if r.Method == stdHttp.MethodOptions {
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
