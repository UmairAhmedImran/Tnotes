package models

import "time"

type Users struct {
	ID        string    `db:"user_id"`
	Username  string    `db:"username"`
	SyncedAt  *time.Time `db:"synced_at"`
	CreatedAt time.Time `db:"created_at"`
}

type Username string

const (
	System Username = "system"
	User   Username = "user"
	Cron   Username = "cron"
);

func (u Username) Valid() bool {
	switch (u) {
		case System, Cron, User:
			return true
	}
	return false
}