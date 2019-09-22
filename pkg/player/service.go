package player

import "github.com/google/uuid"

// Service is the backing player service invoked by HTTP/REST handlers
type Service struct{}

// CreateService creates an instance of the player service struct
func CreateService() *Service {
	return &Service{}
}

// CreatePlayer creates a new player
func (s *Service) CreatePlayer(email string) (IModel, error) {
	return &Model{
		ID:    -1,
		UUID:  uuid.New(),
		Email: email,
	}, nil
}
