package config

import (
	"os"
	"time"
)

type Config struct {
	Bolt   BoltConfig
	Sqlite SQLiteConfig
}

type BoltConfig struct {
	Path     string
	FileMode os.FileMode
}

type SQLiteConfig struct {
	MaxOpenDbConn int
	MaxIdleDbConn int
	MaxDbLifetime time.Duration
}

func Default() Config {
	return Config{
		Bolt: BoltConfig{
			Path:     ".notes/notes.db",
			FileMode: 0600,
		},
		Sqlite: SQLiteConfig{
			MaxOpenDbConn: 1,
			MaxIdleDbConn: 1,
			MaxDbLifetime: 1 * time.Hour,
		},
	}
}
