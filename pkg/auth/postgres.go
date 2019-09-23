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

// CreateAuth creates a new row with uuid and password hash
func (r *PostgresRepository) CreateAuth(uuid uuid.UUID, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	authInsertQuery := `
		INSERT INTO auth (user_uuid, password)
		VALUES ($1, $2)`
	_, err = r.psql.Exec(authInsertQuery, uuid, passwordHash)
	if err != nil {
		return err
	}

	return nil
}
