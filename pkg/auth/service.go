package auth

import (
	"errors"
	"time"

	"github.com/aaclee/mkn-api/pkg/user"
	"github.com/dgrijalva/jwt-go"
)

type userRepository interface {
	FindUserByEmail(email string) (user.IModel, error)
}

// Service is the backing auth service invoked by HTTP/REST handlers
type Service struct {
	user userRepository
}

// CreateService creates an instance of the auth service struct
func CreateService(ur userRepository) *Service {
	return &Service{
		user: ur,
	}
}

// Authenticate validates the username against the password and returns a JWT
func (s *Service) Authenticate(username string, password string) (string, error) {
	user, err := s.user.FindUserByEmail(username)
	// TODO: err from FindUserByEmail needs to be handled
	if err != nil {
		return "", errors.New("Invalid username or password")
	}

	if user.GetEmail() != "admin.mkn@gmail.com" || password != "password" {
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
