package postgres

import (
	"database/sql"
	"fmt"
	"time"
)

// Config represents the requirements to open a postgres connection
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
}

const (
	maxAttempt    = 3
	sleepInSecond = 15
)

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

	var db *sql.DB
	var err error

	openAttempt := 0
	for openAttempt <= maxAttempt {
		openAttempt++

		db, err = sql.Open("postgres", psqlInfo)
		if err == nil {
			break
		}

		fmt.Printf("Attempt %v to open db connection, error: %s; Try again in %d second(s)\n", openAttempt, err, sleepInSecond)
		time.Sleep(sleepInSecond * time.Second)
	}
	if err != nil {
		return nil, err
	}

	pingAttempt := 0
	for pingAttempt <= maxAttempt {
		pingAttempt++

		err = db.Ping()
		if err == nil {
			break
		}

		fmt.Printf("Attempt %v to ping db, error: %s; Try again in %d second(s)\n", pingAttempt, err, sleepInSecond)
		time.Sleep(sleepInSecond * time.Second)
	}
	if err != nil {
		return nil, err
	}

	return db, nil
}
