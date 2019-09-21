package auth

import (
	"errors"
	"time"

	"github.com/aaclee/mkn-api/pkg/player"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type authRepository interface {
	Authenticate(uuid uuid.UUID, password string) error
}

type playerRepository interface {
	FindPlayerByEmail(email string) (player.IModel, error)
}

// Service is the backing auth service invoked by HTTP/REST handlers
type Service struct {
	auth   authRepository
	player playerRepository
}

// CreateService creates an instance of the auth service struct
func CreateService(ar authRepository, pr playerRepository) *Service {
	return &Service{
		auth:   ar,
		player: pr,
	}
}

// Authenticate validates the username against the password and returns a JWT
func (s *Service) Authenticate(username string, password string) (string, error) {
	user, err := s.player.FindPlayerByEmail(username)
	// TODO: err from FindPlayerByEmail needs to be handled
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
