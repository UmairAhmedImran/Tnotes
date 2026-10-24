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
		slog.Error("Error Pinging the database: ", "error", err)
		return err
	}

	slog.Info(" Pinged database successfully ")
	return nil

}

func (db *SqliteDB) Open(dsn string) (*sql.DB, error) {
	var err error

	db.lock.Do(func() {

		db.instance, err = sql.Open("sqlite", dsn)

		if err != nil {
			slog.Error("Error opening the database connection",
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
		slog.Error("SqliteDB is not initialized")
	}
	return db.instance
}
