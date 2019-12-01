package campaign

import (
	"github.com/aaclee/mkn-api/pkg/player"
	"github.com/google/uuid"
)

type campaignRepository interface {
	CreateCampaign(email string) (IModel, error)
	FindCampaignByID(id string) (IModel, error)
	FindCampaignsByPlayerID(id int) ([]IModel, error)
}

type playerRepository interface {
	FindPlayerByUUID(uuid uuid.UUID) (player.IModel, error)
}

// Service is the backing campaign service invoked by HTTP/REST handlers
type Service struct {
	campaign campaignRepository
	player   playerRepository
}

// CreateService creates an instance of the campaign service struct
func CreateService(cr campaignRepository, pr playerRepository) *Service {
	return &Service{
		campaign: cr,
		player:   pr,
	}
}

// CreateCampaign creates a new campaign
func (s *Service) CreateCampaign(name string) (IModel, error) {
	campaign, err := s.campaign.CreateCampaign(name)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}

// FindCampaignByID finds a campaign by :id
func (s *Service) FindCampaignByID(id string) (IModel, error) {
	campaign, err := s.campaign.FindCampaignByID(id)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}

// FindCampaignsByUUID returns all characters from a player by uuid
func (s *Service) FindCampaignsByUUID(playerUUID string) ([]IModel, error) {
	uuid, err := uuid.Parse(playerUUID)
	if err != nil {
		return nil, err
	}

	player, err := s.player.FindPlayerByUUID(uuid)
	if err != nil {
		return nil, err
	}

	campaigns, err := s.campaign.FindCampaignsByPlayerID(player.GetID())
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}
