package campaign

// Model represents the collection of campaign information
type Model struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetName returns the campaign name
func (m *Model) GetName() string {
	return m.Name
}
