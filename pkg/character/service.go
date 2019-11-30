package character

type characterRepository interface {
	CreateCharacter(c *Model) (IModel, error)
}

// Service is the backing character service invoked by HTTP/REST handlers
type Service struct {
	character characterRepository
}

// CreateService creates an instance of the character service struct
func CreateService(cr characterRepository) *Service {
	return &Service{
		character: cr,
	}
}

// CreateCharacter creates a new character
func (s *Service) CreateCharacter(c *Model) (IModel, error) {
	character, err := s.character.CreateCharacter(c)
	if err != nil {
		return nil, err
	}

	return character, nil
}
