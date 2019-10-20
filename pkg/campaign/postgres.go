package campaign

import (
	"database/sql"
)

// IModel represents the model interface
type IModel interface {
	GetName() string
}

// PostgresRepository is the backing campaign repository invoked by services
type PostgresRepository struct {
	psql *sql.DB
}

// CreatePostgresRepository creates an instance of the campaign repository struct
func CreatePostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		psql: db,
	}
}

// CreateCampaign creates a campaign in the postgres database
func (r *PostgresRepository) CreateCampaign(name string) (IModel, error) {
	var campaignID int
	err := r.psql.QueryRow(
		`INSERT INTO campaigns (name)
		VALUES ($1)
		RETURNING id`,
		name,
	).Scan(&campaignID)
	if err != nil {
		return nil, err
	}

	return &Model{
		ID:   campaignID,
		Name: name,
	}, nil
}
