package character

import (
	"errors"

	"github.com/aaclee/mkn-api/pkg/player"
	"github.com/google/uuid"
)

type characterRepository interface {
	CreateCharacter(c *Model) (IModel, error)
	FindCharacterByID(id string) (IModel, error)
	FindCharactersByPlayerID(id int) ([]IModel, error)
	UpdateCharacterByID(c IModel, data map[string]interface{}) (IModel, error)
}

type playerRepository interface {
	FindPlayerByID(id int) (player.IModel, error)
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

// UpdateCharacterByID updates character by id and returns updated character
func (s *Service) UpdateCharacterByID(playerUUID string, character IModel, data map[string]interface{}) (IModel, error) {
	uuid, err := uuid.Parse(playerUUID)
	if err != nil {
		return nil, err
	}

	player, err := s.player.FindPlayerByUUID(uuid)
	if err != nil {
		return nil, err
	}

	if player.GetID() != character.GetPlayerID() {
		return nil, errors.New("character does not belong to player")
	}

	c, err := s.character.UpdateCharacterByID(character, data)
	if err != nil {
		return nil, err
	}

	return c, nil
}
