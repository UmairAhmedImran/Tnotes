package models

import "time"

type Note struct {
	ID             string    `db:"note_id"`
	Title          string    `db:"title"`
	Content        string    `db:"content"`
	WorkspaceID    *int64    `db:"workspace_id"`
	IsPrivate      bool      `db:"is_private"`
	IsDeleted      bool      `db:"is_deleted"`
	SortOrder      int       `db:"sort_order"`
	Source         *string   `db:"source"`
	AdditionalInfo string    `db:"additional_info"` // JSON text
	AddedByID      string    `db:"added_by_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`

	Tags []Tag `db:"-"` // loaded via notes_tags
}

type NoteVersion struct {
	ID           string     `db:"note_version_id"`
	NoteID       string     `db:"note_id"`
	NoteSnapshot string     `db:"note_snapshot"` // JSON text
	SyncedAt     *time.Time `db:"synced_at"`
	CreatedAt    time.Time  `db:"created_at"`
}

type NoteTag struct {
	ID        string    `db:"note_tag_id"`
	NoteID    string    `db:"note_id"`
	TagID     string    `db:"tag_id"`
	CreatedAt time.Time `db:"created_at"`
}