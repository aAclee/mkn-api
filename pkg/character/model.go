package character

// Model represents the collection of campaign information
type Model struct {
	ID         int    `json:"id"`
	PlayerID   int    `json:"playerId"`
	CampaignID int    `json:"campaignId"`
	Name       string `json:"name"`
	FamilyName string `json:"familyName"`
}

// GetName returns the campaign name
func (m *Model) GetName() string {
	return m.Name
}
