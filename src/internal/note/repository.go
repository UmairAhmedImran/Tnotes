package note

import (
	"context"
	"database/sql"
	"log/slog"
	"tnotes/internal/db/models"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Insert(ctx context.Context, data CreateNote, userId string) {

	noteId := uuid.NewString()

	query := `
		INSERT INTO notes(note_id, title, content, added_by_id)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query, noteId, data.Title, data.Content, userId)

	if err != nil {
		slog.Error("[NOTE][REPOSITORY]: Cannot insert user into the DB: ", "error", err)
	}

	slog.Info("[NOTE][REPOSITORY]: Note is saved into the DB: ", "result", result)
}

func (r *Repository) List(ctx context.Context) ([]models.Note, error) {
	query := `
		SELECT note_id, title, content, created_at FROM notes
	`

	var notes []models.Note

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		slog.Error("[NOTE][REPOSITORY]: error in fetching notes list: ", "error", err)
		return notes, err
	}

	defer rows.Close()

	for rows.Next() {
		var note models.Note

		if err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt); err != nil {
			slog.Error("[NOTE][REPOSITORY]: error in fetching notes list: ", "error", err)
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}
