package campaign

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aaclee/mkn-api/pkg/encode"
	"github.com/aaclee/mkn-api/pkg/jwt"
)

type campaignService interface {
	CreateCampaign(email string) (IModel, error)
	FindCampaignByID(id string) (IModel, error)
	FindCampaignsByUUID(uuid string) ([]IModel, error)
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

// FindCampaignByID finds a campaign by :id
func (h *Handler) FindCampaignByID(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.String(), "/")

	ok := len(path) == 4 && path[2] == "campaigns"
	if !ok {
		encode.ErrorJSON(w, http.StatusBadRequest, "ID parameter not found")
		return
	}

	id := path[3]
	campaign, err := h.campaign.FindCampaignByID(id)
	if err != nil {
		encode.ErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	encode.JSON(w, campaign, http.StatusOK)
}

// FindCampaignsByUUID returns all campaigns for the current player by uuid
func (h *Handler) FindCampaignsByUUID(w http.ResponseWriter, r *http.Request) {
	claims, err := jwt.ParseRequest(r)
	if err != nil {
		encode.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Corrupted token: %v", err))
		return
	}

	playerUUID, ok := claims["sub"].(string)
	if !ok {
		encode.ErrorJSON(w, http.StatusBadRequest, "UUID missing from claims")
		return
	}

	campaigns, err := h.campaign.FindCampaignsByUUID(playerUUID)
	if err != nil {
		encode.ErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	encode.JSON(w, campaigns, http.StatusOK)
}
