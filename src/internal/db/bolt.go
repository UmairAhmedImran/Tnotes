package db

import (
	"log/slog"
	"sync"
	"tnotes/internal/config"

	bolt "github.com/boltdb/bolt"
)

type BoltDB struct {
	instance  *bolt.DB
	lock      sync.Once
	initError error
	config    config.BoltConfig
}

var Bolt = BoltDB{
	config: config.Default().Bolt,
}

func (db *BoltDB) Open() (*bolt.DB, error) {

	db.lock.Do(func() {
		db.instance, db.initError = bolt.Open(
			db.config.Path,
			db.config.FileMode,
			nil,
		)

		if db.initError != nil {
			slog.Error("Failed to open Bolt DB",
				"error", db.initError,
			)
		}
	})

	return db.instance, db.initError
}

func (db *BoltDB) Get() (*bolt.DB, error) {
	if db.instance == nil {
		slog.Error("BoltDB is not initialized")
	}
	return db.instance, db.initError
}

func (db *BoltDB) Close() error {
	if db.instance == nil {
		slog.Info("No BoltDB to Close")
		return nil
	}

	err := db.instance.Close()

	if err != nil {
		slog.Warn("Issues in closing the Bolt DB",
			"error", err,
		)
		return err
	}

	slog.Info("BoltDB Closed Successfully")

	return nil

}
