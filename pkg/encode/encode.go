package encode

import (
	"encoding/json"
	"net/http"
)

// JSON writes to a ResponseWriter in json format
func JSON(w http.ResponseWriter, resp interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

// ErrorJSON writes to a ResponseWriter in standardized json format
func ErrorJSON(w http.ResponseWriter, status int, msg string) {
	JSON(w, struct {
		Error string `json:"error"`
	}{
		Error: msg,
	}, status)
}
