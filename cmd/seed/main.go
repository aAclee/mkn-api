package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/aaclee/mkn-api/pkg/logger"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	host = "localhost"
	port = 5432
	user = "mkn_psql"
	pass = "password"
	name = "mkn_db"

	adminEmail    = "admin.mkn@gmail.com"
	adminPassword = "password"
)

func main() {
	// Setting up logger
	logger := logger.CreateLogger()

	// Creating psql connection info
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		pass,
		name,
	)

	// Connecting to psql
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Errorf("SEED: Could not open connection to database: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Errorf("SEED: Could not ping database: %v", err)
		os.Exit(1)
	}

	adminUUID := uuid.New()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Errorf("SEED: Could not create password hash: %v", err)
		os.Exit(1)
	}

	authInsertQuery := `
		INSERT INTO auth (user_uuid, password)
		VALUES ($1, $2)`
	_, err = db.Exec(authInsertQuery, adminUUID, passwordHash)
	if err != nil {
		logger.Errorf("SEED: Could not insert admin auth value: %v", err)
		os.Exit(1)
	}

	userInsertQuery := `
		INSERT INTO players (uuid, email, admin)
		VALUES ($1, $2, $3)`
	_, err = db.Exec(userInsertQuery, adminUUID, adminEmail, true)
	if err != nil {
		logger.Errorf("SEED: Could not insert admin player value: %v", err)
		os.Exit(1)
	}

	logger.Info("Successfully seeded admin in database")
}
