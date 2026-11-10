/*
Copyright © 2026 Tnotes
*/
package main

import (
	// "log"
	// "os"
	"log/slog"
	"tnotes/internal/cli"
	// "tnotes/internal/db"
	// "github.com/joho/godotenv"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelInfo)
	// if err := godotenv.Load(); err != nil {
	// 	log.Printf("Could not load .env: %v", err)
	// }

	// dsn := os.Getenv("DSN")
	// wd, _ := os.Getwd()

	// log.Printf("Working directory: %s", wd)
	// log.Printf("DSN: %q", dsn)

	// dbInstance, err := db.Sqlite.Open(dsn)
	// if err != nil {
	// 	log.Fatalf("Error initializing SQLite DB: %v", err)
	// }
	// defer dbInstance.Close()

	cli.Execute()
}
