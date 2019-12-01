package main

import (
	"flag"
	stdHttp "net/http"
	"os"

	"github.com/aaclee/mkn-api/pkg/campaign"
	"github.com/aaclee/mkn-api/pkg/character"
	"github.com/aaclee/mkn-api/pkg/encode"

	"github.com/aaclee/mkn-api/pkg/auth"
	"github.com/aaclee/mkn-api/pkg/http"
	"github.com/aaclee/mkn-api/pkg/jwt"
	"github.com/aaclee/mkn-api/pkg/logger"
	"github.com/aaclee/mkn-api/pkg/middleware"
	"github.com/aaclee/mkn-api/pkg/player"
	"github.com/aaclee/mkn-api/pkg/postgres"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

const (
	serverConfig       = "./config/server.development.json"
	postgresDevConfig  = "./db/postgres.development.json"
	postgresTestConfig = "./db/postgres.development.json"
)

func main() {
	// Setting up logger
	log := logger.CreateLogger()

	// Getting server configs
	c, err := http.GetServerConfigs(serverConfig)
	if err != nil {
		log.Fatalf("Error processing server config file at: %s, %s\n", serverConfig, err)
	}

	// Getting database configs
	postgresConfig := postgresDevConfig
	if os.Getenv("MODE") == "test" {
		postgresConfig = postgresTestConfig
	}

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
		Log:  log,
		Port: c.Port,
	}

	// Create server
	server := http.CreateServer(config)

	// Create server mux
	r := mux.NewRouter()

	// Mux middleware
	r.Use(standardizeHandler)
	r.Use(logger.Middleware)

	// Repositories
	authRepository := auth.CreatePostgresRepository(db)
	campaignRepository := campaign.CreatePostgresRepository(db)
	characterRepository := character.CreatePostgresRepository(db)
	playerRepository := player.CreatePostgresRepository(db)

	// Services
	authService := auth.CreateService(authRepository, playerRepository)
	campaignService := campaign.CreateService(campaignRepository, playerRepository)
	characterService := character.CreateService(characterRepository, playerRepository)
	playerService := player.CreateService(playerRepository)

	// Handlers
	authHandler := auth.CreateHandler(authService)
	r.HandleFunc("/api/auth", authHandler.Authenticate).Methods("POST")
	r.HandleFunc("/api/auth/confirm", authHandler.ConfirmPlayer).Methods("POST")

	campaignHandler := campaign.CreateHandler(campaignService)
	r.HandleFunc("/api/campaigns", middleware.HandlerFunc(
		campaignHandler.CreateCampaign,
		jwt.MiddlewareVerify,
	)).Methods("POST")
	r.HandleFunc("/api/campaigns/{id}", middleware.HandlerFunc(
		campaignHandler.FindCampaignByID,
		jwt.MiddlewareVerify,
	)).Methods("GET")
	r.HandleFunc("/api/campaigns", middleware.HandlerFunc(
		campaignHandler.FindCampaignsByUUID,
		jwt.MiddlewareVerify,
	)).Methods("GET")

	characterHandler := character.CreateHandler(characterService)
	r.HandleFunc("/api/characters", middleware.HandlerFunc(
		characterHandler.CreateCharacter,
		jwt.MiddlewareVerify,
	)).Methods("POST")
	r.HandleFunc("/api/characters/{id}", middleware.HandlerFunc(
		characterHandler.FindCharacterByID,
		jwt.MiddlewareVerify,
	)).Methods("GET")
	r.HandleFunc("/api/characters", middleware.HandlerFunc(
		characterHandler.FindCharactersByUUID,
		jwt.MiddlewareVerify,
	)).Methods("GET")
	r.HandleFunc("/api/characters/{id}", middleware.HandlerFunc(
		characterHandler.UpdateCharacterByID,
		jwt.MiddlewareVerify,
	)).Methods("PATCH")

	playerHandler := player.CreateHandler(playerService)
	r.HandleFunc("/api/players", middleware.HandlerFunc(
		playerHandler.CreatePlayer,
		player.MiddlewareAdmin,
		jwt.MiddlewareVerify,
	)).Methods("POST")

	// Handler for non-existing routes
	r.PathPrefix("/").HandlerFunc(catchAllHandler)

	// Listen and Serve
	err = server.ListenAndServe(r)
	if err != nil {
		log.Fatalf("Server could not start on port: %d\n", config.Port)
	}
}

func catchAllHandler(w stdHttp.ResponseWriter, r *stdHttp.Request) {
	encode.ErrorJSON(w, stdHttp.StatusNotFound, "Path not found!")
}

func standardizeHandler(next stdHttp.Handler) stdHttp.Handler {
	return stdHttp.HandlerFunc(func(w stdHttp.ResponseWriter, r *stdHttp.Request) {
		if r.Method == stdHttp.MethodOptions {
			return
		}

		w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")

		next.ServeHTTP(w, r)
	})
}
