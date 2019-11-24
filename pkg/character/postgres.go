package character

import (
	"database/sql"
	"fmt"
)

// IModel represents the model interface
type IModel interface {
	GetName() string
}

// PostgresRepository is the backing character repository invoked by services
type PostgresRepository struct {
	psql *sql.DB
}

// CreatePostgresRepository creates an instance of the character repository struct
func CreatePostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		psql: db,
	}
}

// CreateCharacter creates a campaign in the postgres database
func (r *PostgresRepository) CreateCharacter() (IModel, error) {
	fmt.Println("CreateCharacter")
	return &Model{}, nil
}
