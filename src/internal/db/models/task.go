package models

import "time"

type Task struct {
	ID                     string     `db:"task_id"`
	Title                  string     `db:"title"`
	Content                string     `db:"content"`
	DueDate                *time.Time `db:"due_date"`
	Status                 *string    `db:"status"`
	Priority               *string    `db:"priority"`
	WorkspaceID            *int64     `db:"workspace_id"`
	IsPrivate              bool       `db:"is_private"`
	IsDeleted              bool       `db:"is_deleted"`
	SortOrder              int        `db:"sort_order"`
	Source                 *string    `db:"source"`
	AdditionalInfo         string     `db:"additional_info"` // JSON text
	AddedByID              string     `db:"added_by_id"`
	RecurringTaskID        *string    `db:"recurring_task_id"`
	RecurringPeriod        *string    `db:"recurring_period"`
	LastRecurringCreatedAt *time.Time `db:"last_recurring_created_at"`
	CreatedAt              time.Time  `db:"created_at"`
	UpdatedAt              time.Time  `db:"updated_at"`

	Tags []Tag `db:"-"` // loaded via tasks_tags
}

type TaskVersion struct {
	ID           string     `db:"task_version_id"`
	TaskID       string     `db:"task_id"`
	TaskSnapshot string     `db:"task_snapshot"` // JSON text
	SyncedAt     *time.Time `db:"synced_at"`
	CreatedAt    time.Time  `db:"created_at"`
}

type TaskTag struct {
	ID        string    `db:"task_tag_id"`
	TaskID    string    `db:"task_id"`
	TagID     string    `db:"tag_id"`
	CreatedAt time.Time `db:"created_at"`
}