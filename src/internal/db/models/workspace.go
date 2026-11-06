package models

import "time"

type Workspace struct {
	ID            int64     `db:"workspace_id"`
	DirectoryPath string    `db:"directory_path"`
	CreatedAt     time.Time `db:"created_at"`
}
