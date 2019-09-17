package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler is a RESTful HTTP endpoint for for authentication
type Handler struct{}

// CreateHandler creates a new auth handler instance
func CreateHandler() Handler {
	return Handler{}
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

	if body.Username != "admin.mkn@gmail.com" || body.Password != "password" {
		encodeError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	w.WriteHeader(200)
	encode(w, struct {
		Token string `json:"token"`
	}{"abcd-1234-efg-567"})
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
