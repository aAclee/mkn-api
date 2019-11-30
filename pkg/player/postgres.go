package player

import (
	"database/sql"

	"github.com/google/uuid"
)

// IModel represents the model interface
type IModel interface {
	GetEmail() string
	GetID() int
	GetUUID() uuid.UUID
	IsAdmin() bool
}

// PostgresRepository is the backing player repository invoked by services
type PostgresRepository struct {
	psql *sql.DB
}

// CreatePostgresRepository creates an instance of the player repository struct
func CreatePostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		psql: db,
	}
}

// FindPlayerByEmail finds a player by email
func (r *PostgresRepository) FindPlayerByEmail(email string) (IModel, error) {
	row := r.psql.QueryRow(
		`SELECT id, uuid, email, admin FROM players
		WHERE email = $1`,
		email,
	)

	model := &Model{}
	err := row.Scan(&model.ID, &model.UUID, &model.Email, &model.Admin)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// FindPlayerByUUID finds a player by uuid
func (r *PostgresRepository) FindPlayerByUUID(uuid uuid.UUID) (IModel, error) {
	row := r.psql.QueryRow(
		`SELECT id, uuid, email, admin FROM players
		WHERE uuid = $1`,
		uuid,
	)

	model := &Model{}
	err := row.Scan(&model.ID, &model.UUID, &model.Email, &model.Admin)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// CreatePlayer creates a player in the postgres database
func (r *PostgresRepository) CreatePlayer(email string) (IModel, error) {
	var playerID int
	newUUID := uuid.New()
	err := r.psql.QueryRow(
		`INSERT INTO players (uuid, email, admin)
		VALUES ($1, $2, $3)
		RETURNING id`,
		newUUID,
		email,
		false,
	).Scan(&playerID)
	if err != nil {
		return nil, err
	}

	return &Model{
		ID:    playerID,
		UUID:  newUUID,
		Email: email,
	}, nil
}
