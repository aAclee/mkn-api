package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aaclee/mkn-api/pkg/player"
	"github.com/google/uuid"

	"github.com/aaclee/mkn-api/pkg/encode"
)

type authService interface {
	Authenticate(username string, password string) (string, error)
	CreateAuth(uuid uuid.UUID, password string) error
}

type playerService interface {
	FindPlayerByConfirmation(code string) (player.IModel, error)
}

// Handler is a RESTful HTTP endpoint for for authentication
type Handler struct {
	auth   authService
	player playerService
}

// CreateHandler creates a new auth handler instance
func CreateHandler(as authService, ps playerService) *Handler {
	return &Handler{
		auth:   as,
		player: ps,
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

// ConfirmPlayer is the HTTP POST handler for /api/auth/confirm
func (h *Handler) ConfirmPlayer(w http.ResponseWriter, r *http.Request) {
	codes, ok := r.URL.Query()["code"]
	if !ok || len(codes) < 1 {
		encode.ErrorJSON(w, http.StatusBadRequest, "URL param 'code' is required")
		return
	}

	code := codes[0]
	player, err := h.player.FindPlayerByConfirmation(code)

	decoder := json.NewDecoder(r.Body)

	var body struct {
		Username  string `json:"username"`
		PasswordA string `json:"password"`
		PasswordB string `json:"passwordConfirmation"`
	}

	err = decoder.Decode(&body)
	if err != nil {
		encode.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request: %v", err))
		return
	}

	if player.GetEmail() != body.Username {
		encode.ErrorJSON(w, http.StatusBadRequest, "Error parsing request: username does not match")
		return
	}

	if player.GetUUID().String() != code {
		encode.ErrorJSON(w, http.StatusBadRequest, "Error parsing request: confirmation code is invalid")
		return
	}

	if body.PasswordA != body.PasswordB {
		encode.ErrorJSON(w, http.StatusBadRequest, "Error parsing request: password does not match confirmation password")
		return
	}

	err = h.auth.CreateAuth(player.GetUUID(), body.PasswordA)
	if err != nil {
		encode.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Error confirming player: %v", err))
		return
	}

	encode.JSON(w, player, http.StatusCreated)
}
