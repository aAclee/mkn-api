package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aaclee/mkn-api/pkg/encode"
)

type authService interface {
	Authenticate(username string, password string) (string, error)
}

// Handler is a RESTful HTTP endpoint for for authentication
type Handler struct {
	auth authService
}

// CreateHandler creates a new auth handler instance
func CreateHandler(as authService) *Handler {
	return &Handler{
		auth: as,
	}
}

// Authenticate is the HTTP POST handler for /api/auth
func (h *Handler) Authenticate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := decoder.Decode(&body)
	if err != nil {
		encode.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request: %v", err))
		return
	}

	token, err := h.auth.Authenticate(body.Username, body.Password)
	if err != nil {
		encode.ErrorJSON(w, http.StatusUnauthorized, err.Error())
		return
	}

	encode.JSON(w, struct {
		Token string `json:"token"`
	}{token}, http.StatusOK)
}
