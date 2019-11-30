package character

import (
	"github.com/aaclee/mkn-api/pkg/player"
	"github.com/google/uuid"
)

type characterRepository interface {
	CreateCharacter(c *Model) (IModel, error)
	FindCharacterByID(id string) (IModel, error)
	FindCharactersByPlayerID(id int) ([]IModel, error)
}

type playerRepository interface {
	FindPlayerByUUID(uuid uuid.UUID) (player.IModel, error)
}

// Service is the backing character service invoked by HTTP/REST handlers
type Service struct {
	character characterRepository
	player    playerRepository
}

// CreateService creates an instance of the character service struct
func CreateService(cr characterRepository, pr playerRepository) *Service {
	return &Service{
		character: cr,
		player:    pr,
	}
}

// CreateCharacter creates a new character
func (s *Service) CreateCharacter(c *Model, playerUUID string) (IModel, error) {
	uuid, err := uuid.Parse(playerUUID)
	if err != nil {
		return nil, err
	}

	player, err := s.player.FindPlayerByUUID(uuid)
	if err != nil {
		return nil, err
	}

	c.PlayerID = player.GetID()
	character, err := s.character.CreateCharacter(c)
	if err != nil {
		return nil, err
	}

	return character, nil
}

// FindCharacterByID finds a character by :id
func (s *Service) FindCharacterByID(id string) (IModel, error) {
	character, err := s.character.FindCharacterByID(id)
	if err != nil {
		return nil, err
	}

	return character, nil
}

// FindCharactersByUUID returns all characters from a player by uuid
func (s *Service) FindCharactersByUUID(playerUUID string) ([]IModel, error) {
	uuid, err := uuid.Parse(playerUUID)
	if err != nil {
		return nil, err
	}

	player, err := s.player.FindPlayerByUUID(uuid)
	if err != nil {
		return nil, err
	}

	characters, err := s.character.FindCharactersByPlayerID(player.GetID())
	if err != nil {
		return nil, err
	}

	return characters, nil
}
