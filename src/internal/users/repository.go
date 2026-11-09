package users

import (
	"context"
	"database/sql"
	"log/slog"
	"tnotes/internal/db/models"
)

const userCols = "user_id, username, synced_at, created_at"

type Repository struct {
	db	*sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) RetrieveUser(ctx context.Context, username models.Username) (models.Users, error) {
	query := `
		SELECT ? FROM users
		WHERE username = ?
	`

	var user models.Users

	row := r.db.QueryRowContext(ctx, query, userCols, username)

	err := row.Scan(
		&user.ID, &user.Username, &user.SyncedAt, &user.CreatedAt,
	)

	if err != nil {
		slog.Error("Error scanning the user table ", "error", err )
	}

	return user, err
}