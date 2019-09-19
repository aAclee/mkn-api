package auth

import (
	"errors"
	"time"

	"github.com/aaclee/mkn-api/pkg/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type authRepository interface {
	Authenticate(uuid uuid.UUID, password string) error
}

type userRepository interface {
	FindUserByEmail(email string) (user.IModel, error)
}

// Service is the backing auth service invoked by HTTP/REST handlers
type Service struct {
	auth authRepository
	user userRepository
}

// CreateService creates an instance of the auth service struct
func CreateService(ar authRepository, ur userRepository) *Service {
	return &Service{
		auth: ar,
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

	err = s.auth.Authenticate(user.GetUUID(), password)
	if err != nil {
		return "", errors.New("Invalid username or password")
	}

	// TODO: Add email / username to this claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "munchkin-api",
		"sub": user.GetUUID(),
		"usr": user.GetEmail(),
		"aud": "munchkin-ui",
		"exp": time.Now().Add(time.Hour * 8),
		"nbf": time.Now(),
		"iat": time.Now(),
	})

	// TODO: Issue 500 error if token signing failed
	tokenString, _ := token.SignedString([]byte("munchkin-secret"))

	return tokenString, nil
}
