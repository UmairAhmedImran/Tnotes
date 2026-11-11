package users

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"tnotes/internal/db/models"
)

const userCols = "user_id, username, synced_at, created_at"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) RetrieveUser(ctx context.Context, username models.Username) (models.Users, error) {
	slog.Debug("[USER][REPOSITORY]: ", "username", username)

	query := `SELECT ` + userCols + ` FROM users WHERE username = ?`

	slog.Debug("[USER][REPOSITORY]: retrieve user", "query", query, "arg", username)

	var user models.Users

	row := r.db.QueryRowContext(ctx, query, username)

	err := row.Scan(
		&user.ID, &user.Username, &user.SyncedAt, &user.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		slog.Error("[USER][REPOSITORY]: not found", "error", err)
		return user, sql.ErrNoRows
	}

	if err != nil {
		slog.Error("[USER][REPOSITORY]: Error scanning the user table ", "error", err)
	}

	return user, err
}
