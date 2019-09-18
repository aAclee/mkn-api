package postgres

import (
	"database/sql"
	"fmt"
)

// Config represents the requirements to open a postgres connection
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
}

// CreateConnection creates a postgres database connection
func CreateConnection(config Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
