package character

/**
 * This package, specifically UpdateCharacterByID is why we need to use an ORM
 */

import (
	"database/sql"
)

// IModel represents the model interface
type IModel interface {
	GetID() int
	GetName() string
	GetPlayerID() int
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

// FindCharacterByID finds a campaign by :id
func (r *PostgresRepository) FindCharacterByID(id string) (IModel, error) {
	row := r.psql.QueryRow(
		`SELECT id, player_id, campaign_id, name, family_name FROM characters_basic
		WHERE id = $1`,
		id,
	)

	model := &Model{}
	err := row.Scan(
		&model.ID,
		&model.PlayerID,
		&model.CampaignID,
		&model.Name,
		&model.FamilyName,
	)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// FindCharactersByPlayerID returns all characters from a player by id
func (r *PostgresRepository) FindCharactersByPlayerID(playerID int) ([]IModel, error) {
	var models []IModel
	rows, err := r.psql.Query(
		`SELECT id, player_id, campaign_id, name, family_name FROM characters_basic
		WHERE player_id = $1`,
		playerID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		model := &Model{}
		err := rows.Scan(
			&model.ID,
			&model.PlayerID,
			&model.CampaignID,
			&model.Name,
			&model.FamilyName,
		)
		if err != nil {
			return nil, err
		}

		models = append(models, model)
	}

	return models, nil
}

// UpdateCharacterByID returns updated character found by id
func (r *PostgresRepository) UpdateCharacterByID(c IModel, data map[string]interface{}) (IModel, error) {
	row := r.psql.QueryRow(
		`SELECT id, player_id, campaign_id, name, family_name FROM characters_basic
		WHERE id = $1`,
		c.GetID(),
	)

	character := &Model{}
	err := row.Scan(
		&character.ID,
		&character.PlayerID,
		&character.CampaignID,
		&character.Name,
		&character.FamilyName,
	)
	if err != nil {
		return nil, err
	}

	campaignID, ok := data["campaignId"]
	if ok {
		cid, ok := campaignID.(float64)
		if ok {
			character.CampaignID.Valid = true
			character.CampaignID.Int32 = int32(cid)
		} else {
			character.CampaignID.Valid = false
			character.CampaignID.Int32 = 0
		}
	}

	name, ok := data["name"]
	if ok {
		n, ok := name.(string)
		if ok {
			character.Name.Valid = true
			character.Name.String = n
		} else {
			character.Name.Valid = false
			character.Name.String = ""
		}
	}

	familyName, ok := data["familyName"]
	if ok {
		fn, ok := familyName.(string)
		if ok {
			character.FamilyName.Valid = true
			character.FamilyName.String = fn
		} else {
			character.FamilyName.Valid = false
			character.FamilyName.String = ""
		}
	}

	_, err = r.psql.Exec(
		`UPDATE characters_basic
		SET campaign_id=$1, name=$2, family_name=$3
		WHERE id = $4`,
		character.CampaignID,
		character.Name,
		character.FamilyName,
		c.GetID(),
	)

	return character, nil
}
