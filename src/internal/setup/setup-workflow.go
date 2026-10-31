package setup

import (
	"embed"
	"log/slog"
	"tnotes/internal/db"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.up.sql
var embedMigrations embed.FS

func SetupWorkflow() {

	var db, err = db.Sqlite.Open()

	if err != nil {
		slog.Debug("Nothing to do since database does not work correctly")
		return
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		slog.Error("failed to set dialect: ", "error", err)
		return
	}

	goose.SetBaseFS(embedMigrations)

	err = goose.Up(db, "migrations")

	if err != nil {
		slog.Error("Migration could not be applied: ", "error", err)
		return
	}

	slog.Info("Migration applied successful")
}