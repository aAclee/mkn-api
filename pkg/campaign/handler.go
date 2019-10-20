package campaign

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aaclee/mkn-api/pkg/encode"
)

type campaignService interface {
	CreateCampaign(email string) (IModel, error)
}

// Handler is a RESTful HTTP endpoint for for campaigns
type Handler struct {
	campaign campaignService
}

// CreateHandler creates a new campaign handler instance
func CreateHandler(cs campaignService) *Handler {
	return &Handler{
		campaign: cs,
	}
}

// CreateCampaign is the HTTP POST handler for /api/campaigns
func (h *Handler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var body struct {
		Name string `json:"name"`
	}

	err := decoder.Decode(&body)
	if err != nil {
		encode.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request: %v", err))
		return
	}

	if body.Name == "" {
		encode.ErrorJSON(w, http.StatusBadRequest, "Name cannot be empty")
		return
	}

	campaign, err := h.campaign.CreateCampaign(body.Name)
	if err != nil {
		encode.ErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	encode.JSON(w, campaign, http.StatusCreated)
}
