package postgres

import (
	"encoding/json"
	"errors"
	"os"
)

// Configuration type for code-challenge server configs
type Configuration struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"pass"`
	Database string `json:"dbName"`
}

// GetDatabaseConfigs fetches server configs from filenames
func GetDatabaseConfigs(filename string) (*Configuration, error) {
	configuration := Configuration{}

	err := Parse(filename, &configuration)
	if err != nil {
		return nil, err
	}

	err = validateConfig(configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}

func validateConfig(config Configuration) error {
	if config.Host == "" {
		return errors.New("Hostname not provided in config file")
	}

	if config.Port == 0 {
		return errors.New("Port number not provided in config file")
	}

	if config.User == "" {
		return errors.New("Username not provided in config file")
	}

	if config.Password == "" {
		return errors.New("Password not provided in config file")
	}

	if config.Database == "" {
		return errors.New("Database name not provided in config file")
	}

	return nil
}

// Parse parses a json path/filename and fills the configuration
func Parse(filename string, configuration *Configuration) error {
	config, err := os.Open(filename)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(config)
	err = decoder.Decode(&configuration)
	if err != nil {
		return err
	}

	return nil
}
