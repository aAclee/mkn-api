package player

import "github.com/google/uuid"

// Model representats the collection of player information
type Model struct {
	ID    int       `json:"id"`
	UUID  uuid.UUID `json:"uuid"`
	Email string    `json:"email"`
	Admin bool      `json:"admin"`
}

// GetEmail returns the player email
func (m *Model) GetEmail() string {
	return m.Email
}

// GetUUID returns the player email
func (m *Model) GetUUID() uuid.UUID {
	return m.UUID
}

// IsAdmin return admin status of player
func (m *Model) IsAdmin() bool {
	return m.Admin
}
