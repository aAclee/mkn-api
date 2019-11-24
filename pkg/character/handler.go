package character

import (
	"fmt"
	"net/http"

	"github.com/aaclee/mkn-api/pkg/encode"
)

type characterService interface {
	CreateCharacter() (IModel, error)
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
	fmt.Println(r.Body)

	character, err := h.character.CreateCharacter()
	if err != nil {
		encode.ErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	encode.JSON(w, character, http.StatusCreated)
}
