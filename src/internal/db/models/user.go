package models

import "time"

type Users struct {
	ID        string    `db:"user_id"`
	Username  string    `db:"username"`
	SyncedAt  time.Time `db:"synced_at"`
	CreatedAt time.Time `db:"created_at"`
}
