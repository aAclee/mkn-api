package http_test

import (
	"testing"

	"github.com/aaclee/mkn-api/pkg/http"
)

func TestGetServerConfigs(t *testing.T) {
	tests := []struct {
		label    string
		filename string
		port     int
	}{
		{
			label:    "Success",
			filename: "./test_data/config.json",
			port:     8000,
		},
		{
			label:    "Error",
			filename: "./test_data/error.json",
			port:     -1,
		},
	}

	for _, test := range tests {
		config, err := http.GetServerConfigs(test.filename)
		if test.label == "Success" {
			if config.Port != test.port {
				t.Errorf("%s: Port: %v; Expected %v", test.label, config.Port, test.port)
			}
		}

		if test.label == "Error" {
			if err == nil {
				t.Errorf("%s: err is nil, expected a problem", test.label)
			}
		}
	}
}
