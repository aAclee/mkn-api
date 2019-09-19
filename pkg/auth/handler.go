package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
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
		encodeError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request: %v", err))
		return
	}

	token, err := h.auth.Authenticate(body.Username, body.Password)
	if err != nil {
		encodeError(w, http.StatusUnauthorized, err.Error())
		return
	}

	w.WriteHeader(200)
	encode(w, struct {
		Token string `json:"token"`
	}{token})
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
