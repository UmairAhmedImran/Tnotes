/*
Copyright © 2026 Tnotes
*/
package main

import (
	"log"
	"os"

	"UmairAhmedImran/cmd"
	"UmairAhmedImran/internal/db"
	"UmairAhmedImran/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	if _, err := os.Stat(service.NotesDir); !(os.IsNotExist(err)) {
		db.Open()
		defer db.Close()
	}

	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectSql(dsn)
	if err != nil {
		log.Fatal("Error Initializing Sqlite Db")
	}
	defer dbConn.DB.Close()
	cmd.Execute()
}
