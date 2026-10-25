package models

import "time"

type SyncHistory struct {
	ID        string    `db:"sync_history_id"`
	Start     time.Time `db:"start"`
	End       time.Time `db:"end"`
	Metadata  string    `db:"metadata"` // JSON text
	CreatedAt time.Time `db:"created_at"`
}

type Log struct {
	ID        string    `db:"log_id"`
	Action    string    `db:"action"`
	Metadata  string    `db:"metadata"` // JSON text
	Status    *string   `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}

type Provider struct {
	ID            string    `db:"provider_id"`
	Name          string    `db:"name"`
	URL           *string   `db:"url"`
	Configuration string    `db:"configuration"` // JSON text
	CreatedAt     time.Time `db:"created_at"`
}