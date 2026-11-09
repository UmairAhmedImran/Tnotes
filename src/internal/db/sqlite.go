package db

import (
	"database/sql"
	"log/slog"
	"sync"
	"tnotes/internal/config"

	_ "github.com/glebarez/go-sqlite"
)

type SqliteDB struct {
	instance *sql.DB
	lock     sync.Once
	config   config.SQLiteConfig
}

var Sqlite = SqliteDB{
	config: config.Default().Sqlite,
}

func (db *SqliteDB) testConn() error {
	err := db.instance.Ping()

	if err != nil {
		slog.Error("[DB][SQLITE] Error Pinging the database: ", "error", err)
		return err
	}

	slog.Info("[DB][SQLITE] Pinged database successfully ")
	return nil

}

func (db *SqliteDB) Open() (*sql.DB, error) {
	var err error

	db.lock.Do(func() {

		var dsn = db.config.Path + "?_journal_mode=WAL&_busy_timeout=5000&_pragma=foreign_keys(1)"

		db.instance, err = sql.Open("sqlite", dsn)

		if err != nil {
			slog.Error("[DB][SQLITE] Error opening the database connection",
				"error", err,
			)
			return
		}

		err = db.testConn()

		if err != nil {
			return
		}

		db.instance.SetMaxOpenConns(db.config.MaxOpenDbConn)
		db.instance.SetMaxIdleConns(db.config.MaxIdleDbConn)
		db.instance.SetConnMaxLifetime(db.config.MaxDbLifetime)
	})

	return db.instance, err
}

func (db *SqliteDB) Get() *sql.DB {
	if db.instance == nil {
		slog.Error("[DB][SQLITE] SqliteDB is not initialized")
	}
	return db.instance
}
