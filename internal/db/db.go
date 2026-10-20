package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	bolt "github.com/boltdb/bolt"
	_ "github.com/glebarez/go-sqlite"
)

var (
	DbFile   string = ".notes/notes.db"
	instance *bolt.DB
	once     sync.Once
)

type DB struct {
	DB *sql.DB
}

const maxOpenDbConn = 1
const maxIdleDbConn = 1
const maxDbLifetime = 1 * time.Hour

func ConnectSql(dsn string) (*DB, error) {
	d, err := sql.Open("sqlite", dsn)
	if err != nil {
		fmt.Println("Error opening the database connection", err)
		return nil, err
	}
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	err = testDb(d)
	return &DB{DB: d}, nil
}

func testDb(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error", err)
		return err
	}
	return nil
}

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
