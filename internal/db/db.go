package db

import (
	"log"
	"sync"

	bolt "github.com/boltdb/bolt"
)

var (
	DbFile   string = ".notes/notes.db"
	instance *bolt.DB
	once     sync.Once
)

// Open initializes and returns the global DB connection.
func Open() *bolt.DB {
	path := DbFile
	once.Do(func() {
		db, err := bolt.Open(path, 0600, nil)
		if err != nil {
			log.Fatalf("failed to open database: %v", err)
		}
		instance = db
	})
	return instance
}

// Get returns the active DB instance.
func Get() *bolt.DB {
	if instance == nil {
		log.Fatal("database not initialized; call db.Open() first")
	}
	return instance
}

// Close closes the DB when exiting.
func Close() {
	if instance != nil {
		instance.Close()
	}
}
