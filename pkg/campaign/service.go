package campaign

type campaignRepository interface {
	CreateCampaign(email string) (IModel, error)
	FindCampaignByID(id string) (IModel, error)
}

// Service is the backing campaign service invoked by HTTP/REST handlers
type Service struct {
	campaign campaignRepository
}

// CreateService creates an instance of the campaign service struct
func CreateService(cr campaignRepository) *Service {
	return &Service{
		campaign: cr,
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
