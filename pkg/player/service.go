package player

import "github.com/google/uuid"

type playerRepository interface {
	CreatePlayer(email string) (IModel, error)
	FindPlayerByUUID(uuid uuid.UUID) (IModel, error)
}

// Service is the backing player service invoked by HTTP/REST handlers
type Service struct {
	player playerRepository
}

// CreateService creates an instance of the player service struct
func CreateService(pr playerRepository) *Service {
	return &Service{
		player: pr,
	}
}

// CreatePlayer creates a new player
func (s *Service) CreatePlayer(email string) (IModel, error) {
	player, err := s.player.CreatePlayer(email)
	if err != nil {
		return nil, err
	}

	return player, nil
}

// FindPlayerByConfirmation finds a player by confirmation code
func (s *Service) FindPlayerByConfirmation(code string) (IModel, error) {
	uuid, err := uuid.Parse(code)
	if err != nil {
		return nil, err
	}

	player, err := s.player.FindPlayerByUUID(uuid)
	if err != nil {
		return nil, err
	}

	return player, nil
}
