package player

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type playerService interface {
	CreatePlayer(email string) (IModel, error)
}

// Handler is a RESTful HTTP endpoint for for players
type Handler struct {
	player playerService
}

// CreateHandler creates a new auth handler instance
func CreateHandler(ps playerService) *Handler {
	return &Handler{
		player: ps,
	}
}

// CreatePlayer is the HTTP POST handler for /api/players
func (h *Handler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var body struct {
		Email string `json:"email"`
	}

	err := decoder.Decode(&body)
	if err != nil {
		encodeError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request: %v", err))
		return
	}

	if body.Email == "" {
		encodeError(w, http.StatusBadRequest, "Email cannot be empty")
		return
	}

	player, err := h.player.CreatePlayer(body.Email)
	if err != nil {
		encodeError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	encode(w, player)
}

func encode(w http.ResponseWriter, resp interface{}) {
	json.NewEncoder(w).Encode(resp)
}

func encodeError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	encode(w, struct {
		Error string `json:"error"`
	}{
		Error: msg,
	})
}
