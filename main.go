/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"UmairAhmedImran/cmd"
	"UmairAhmedImran/internal/db"
	"UmairAhmedImran/internal/service"
	"log"
	"os"
)

func main() {
	if _, err := os.Stat(service.NotesDir); !(os.IsNotExist(err)) {
		db.Open()
		defer db.Close()
	}

	dsn := "my_database.db?_journal_mode=WAL&_busy_timeout=5000&_pragma=foreign_keys(1)"
	dbConn, err := db.ConnectSql(dsn)
	if err != nil {
		log.Fatal("Error Initializing Sqlite Db")
	}
	defer dbConn.DB.Close()
	cmd.Execute()
}
