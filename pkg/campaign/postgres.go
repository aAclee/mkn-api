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

// FindCampaignByID finds a campaign by :id
func (r *PostgresRepository) FindCampaignByID(id string) (IModel, error) {
	row := r.psql.QueryRow(
		`SELECT id, name FROM campaigns
		WHERE id = $1`,
		id,
	)

	model := &Model{}
	err := row.Scan(&model.ID, &model.Name)
	if err != nil {
		return nil, err
	}

	return model, nil
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

// FindCampaignsByPlayerID returns all campaigns from a player by id
func (r *PostgresRepository) FindCampaignsByPlayerID(playerID int) ([]IModel, error) {
	models := []IModel{}
	rows, err := r.psql.Query(
		`SELECT campaigns.id, campaigns.name FROM campaigns
		JOIN characters_basic ON campaigns.id=characters_basic.campaign_id
		WHERE characters_basic.player_id = $1`,
		playerID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		model := &Model{}
		err := rows.Scan(
			&model.ID,
			&model.Name,
		)
		if err != nil {
			return nil, err
		}

		models = append(models, model)
	}

	return models, nil
}
