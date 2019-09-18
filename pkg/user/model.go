package user

import "github.com/google/uuid"

// Model representats the collection of user information
type Model struct {
	ID    int       `json:"id"`
	UUID  uuid.UUID `json:"uuid"`
	Email string    `json:"email"`
}

// GetEmail returns the user email
func (m *Model) GetEmail() string {
	return m.Email
}
