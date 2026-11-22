package setup

import (
	"embed"
	"encoding/json"
	"log/slog"
	"os"
	"tnotes/internal/db"
	"tnotes/internal/service"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func SetupDBLayer() {

	var db, err = db.Sqlite.Open()

	if err != nil {
		slog.Debug("[SETUP] Nothing to do since database does not work correctly")
		return
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		slog.Error("[SETUP] failed to set dialect: ", "error", err)
		return
	}

	goose.SetBaseFS(embedMigrations)

	err = goose.Up(db, "migrations")

	if err != nil {
		slog.Error("[SETUP] Migration could not be applied: ", "error", err)
		return
	}

	slog.Info("[SETUP] Migration applied successful")
}

func CreateConfigFile() error {
	conFile, err := os.Create(service.ConfigFile)
	if err != nil {
		slog.Error("[CONFIG] Error creating the config file", "error", err)
	}
	defer conFile.Close()

	encoder := json.NewEncoder(conFile)
	err = encoder.Encode(service.BasicConfData)

	if err != nil {
		slog.Error("[CONFIG] Error encoding JSON: ", "error", err)
		os.Exit(1)
	}

	slog.Info("[CONFIG] Created config file")
	return nil
}

func CheckandCreateDirectory() {
	if _, err := os.Stat(service.NotesDir); os.IsNotExist(err) {
		os.Mkdir(service.NotesDir, os.ModePerm)

	} else {
		slog.Info("[DIR] Notes Directory already exists.")
		os.Exit(1)
	}
}

func CallInit() error {
	_, err := os.Getwd()

	if err != nil {
		slog.Error("[INIT] Error getting current directory")
	}

	CheckandCreateDirectory()
	CreateConfigFile()
	SetupDBLayer()

	slog.Info("[INIT] Tnotes Initialization Complete")

	return nil
}
