package note

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db};
}

func (r *Repository) Insert(ctx context.Context, data CreateNote, userId string) {

	noteId := uuid.NewString()

	query := `
		INSERT INTO NOTES(note_id, title, content, added_by_id)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query, noteId, data.Title, data.Content, userId)

	if err != nil {
		slog.Error("Cannot insert user into the DB: ", "error", err)
	}

	slog.Info("Note is saved into the DB: ", "result", result)
}