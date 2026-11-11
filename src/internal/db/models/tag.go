package models

import "time"

type Tag struct {
	ID        string     `db:"tag_id"`
	Title     string     `db:"title"`
	SyncedAt  *time.Time `db:"synced_at"`
	CreatedAt time.Time  `db:"created_at"`
}
