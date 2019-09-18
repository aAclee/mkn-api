package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Service is the backing auth service invoked by HTTP/REST handlers
type Service struct{}

// CreateService creates an instance of the auth service struct
func CreateService() *Service {
	return &Service{}
}

// Authenticate validates the username against the password and returns a JWT
func (h *Service) Authenticate(username string, password string) (string, error) {
	if username != "admin.mkn@gmail.com" || password != "password" {
		return "", errors.New("Invalid username or password")
	}

	// TODO: Add email / username to this claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "munchkin-api",
		"sub": "admin.mkn@gmail.com",
		"usr": "admin.mkn@gmail.com",
		"aud": "munchkin-ui",
		"exp": time.Now().Add(time.Hour * 8),
		"nbf": time.Now(),
		"iat": time.Now(),
	})

	// TODO: Issue 500 error if token signing failed
	tokenString, _ := token.SignedString([]byte("munchkin-secret"))

	return tokenString, nil
}
