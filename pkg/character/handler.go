package character

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aaclee/mkn-api/pkg/encode"
)

type characterService interface {
	CreateCharacter(c *Model) (IModel, error)
	FindCharacterByID(id string) (IModel, error)
}

// Handler is a RESTful HTTP endpoint for for characters
type Handler struct {
	character characterService
}

// CreateHandler creates a new character handler instance
func CreateHandler(cs characterService) *Handler {
	return &Handler{
		character: cs,
	}
}

// CreateCharacter is the HTTP POST handler for /api/characters
func (h *Handler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	c := &Model{}
	err := decoder.Decode(c)
	if err != nil {
		encode.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request: %v", err))
		return
	}

	character, err := h.character.CreateCharacter(c)
	if err != nil {
		encode.ErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	encode.JSON(w, character, http.StatusCreated)
}

// FindCharacterByID finds a character by :id
func (h *Handler) FindCharacterByID(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.String(), "/")

	ok := len(path) == 4 && path[2] == "characters"
	if !ok {
		encode.ErrorJSON(w, http.StatusBadRequest, "ID parameter not found")
		return
	}

	id := path[3]
	character, err := h.character.FindCharacterByID(id)
	if err != nil {
		encode.ErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	encode.JSON(w, character, http.StatusOK)
}
