package player

import (
	"encoding/json"
	"net/http"
)

// Handler is a RESTful HTTP endpoint for for players
type Handler struct{}

// CreateHandler creates a new auth handler instance
func CreateHandler() *Handler {
	return &Handler{}
}

// CreatePlayer is the HTTP POST handler for /api/players
func (h *Handler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	encode(w, struct {
		Message string `json:"message"`
	}{"POST /api/players"})
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
