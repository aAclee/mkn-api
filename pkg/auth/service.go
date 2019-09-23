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
	CreateAuth(uuid uuid.UUID, password string) error
}

type playerRepository interface {
	FindPlayerByEmail(email string) (player.IModel, error)
	FindPlayerByUUID(uuid uuid.UUID) (player.IModel, error)
}

// ConfirmationInfo is the required information to confirm a player
type ConfirmationInfo struct {
	confirmationCode string
	username         string
	password         string
	confirmPassword  string
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

// CreateAuth creates a new entry for authentication
func (s *Service) CreateAuth(ci *ConfirmationInfo) (player.IModel, error) {
	uuid, err := uuid.Parse(ci.confirmationCode)
	if err != nil {
		return nil, err
	}

	player, err := s.player.FindPlayerByUUID(uuid)
	if err != nil {
		return nil, err
	}

	if player.GetEmail() != ci.username {
		return nil, errors.New("username does not match")
	}

	if player.GetUUID().String() != ci.confirmationCode {
		return nil, errors.New("confirmation code is invalid")
	}

	if ci.password != ci.confirmPassword {
		return nil, errors.New("password does not match confirmation password")
	}

	err = s.auth.CreateAuth(uuid, ci.password)
	if err != nil {
		return nil, err
	}

	return player, nil
}
