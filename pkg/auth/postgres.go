package auth

import (
	"database/sql"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// PostgresRepository is the backing auth repository invoked by services
type PostgresRepository struct {
	psql *sql.DB
}

// CreatePostgresRepository creates an instance of the auth repository struct
func CreatePostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		psql: db,
	}
}

// Authenticate validates the uuid against the password hash and returns a JWT
func (r *PostgresRepository) Authenticate(uuid uuid.UUID, password string) error {
	row := r.psql.QueryRow(
		`SELECT password FROM auth
		WHERE user_uuid = $1`,
		uuid,
	)

	var passwordHash string
	err := row.Scan(&passwordHash)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
