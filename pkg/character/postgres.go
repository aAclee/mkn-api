package character

import (
	"database/sql"
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
func (r *PostgresRepository) CreateCharacter(c *Model) (IModel, error) {
	err := r.psql.QueryRow(
		`INSERT INTO characters_basic (player_id, campaign_id, name, family_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id`,
		c.PlayerID,
		c.CampaignID,
		c.Name,
		c.FamilyName,
	).Scan(&c.ID)
	if err != nil {
		return nil, err
	}

	return c, nil
}
