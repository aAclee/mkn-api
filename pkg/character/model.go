package character

import (
	"database/sql"
	"encoding/json"
	"errors"
)

// Model represents the collection of campaign information
type Model struct {
	ID         int            `json:"id"`
	PlayerID   int            `json:"playerId"`
	CampaignID sql.NullInt32  `json:"campaignId"`
	Name       sql.NullString `json:"name"`
	FamilyName sql.NullString `json:"familyName"`
}

// GetName returns the campaign name
func (m *Model) GetName() string {
	return m.Name.String
}

// UnmarshalJSON is a custom unmarshaler used to handle nil/null/undefined values
func (m *Model) UnmarshalJSON(data []byte) error {
	var rm map[string]interface{}
	json.Unmarshal(data, &rm)

	v, ok := rm["playerId"]
	if !ok {
		return errors.New("playerId is required")
	}

	n, ok := v.(float64)
	if !ok {
		return errors.New("playerId must be a number")
	}

	m.PlayerID = int(n)

	m.CampaignID = sql.NullInt32{}
	v, ok = rm["campaignId"]
	if ok {
		if n, ok := v.(int32); ok {
			m.CampaignID.Valid = true
			m.CampaignID.Int32 = n
		}
	}

	m.Name = sql.NullString{}
	v, ok = rm["name"]
	if ok {
		if s, ok := v.(string); ok {
			m.Name.Valid = true
			m.Name.String = s
		}
	}

	m.FamilyName = sql.NullString{}
	v, ok = rm["familyName"]
	if ok {
		if s, ok := v.(string); ok {
			m.FamilyName.Valid = true
			m.FamilyName.String = s
		}
	}

	return nil
}
