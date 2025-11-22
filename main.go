/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"UmairAhmedImran/cmd"
	"UmairAhmedImran/internal/db"
	"UmairAhmedImran/internal/service"
	"os"
)

func main() {
	if _, err := os.Stat(service.NotesDir); !(os.IsNotExist(err)) {
		db.Open()
		defer db.Close()
	}
	cmd.Execute()
}
