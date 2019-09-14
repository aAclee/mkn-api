package http

import (
	"encoding/json"
	"errors"
	"os"
)

// Configuration type for code-challenge server configs
type Configuration struct {
	Port int `json:"port"`
}

// GetServerConfigs fetches server configs from filenames
func GetServerConfigs(filename string) (*Configuration, error) {
	configuration := Configuration{
		Port: -1,
	}

	err := Parse(filename, &configuration)
	if err != nil {
		return nil, err
	}

	if configuration.Port == -1 {
		return nil, errors.New("Port number not provided in config file")
	}

	return &configuration, nil
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
