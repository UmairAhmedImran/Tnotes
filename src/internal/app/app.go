package app

import (
	"database/sql"
	"log/slog"
	"tnotes/internal/db"
	"tnotes/internal/note"
	"tnotes/internal/setup"
	"tnotes/internal/users"
)

type App struct {
	Notes *note.Service
	db    *sql.DB
}

var app *App

func GetApp() (*App, error) {
	if app != nil {
		return app, nil
	}

	var err error

	app = &App{}
	app, err = app.New()
	setup.SetupWorkflow()

	return app, err
}

func (a *App) New() (*App, error) {
	db, err := db.Sqlite.Open()

	if err != nil {
		slog.Error("database not working")
	}

	usersSvc := users.NewService(db)
	noteSvc := note.NewService(db, usersSvc)

	return &App{Notes: noteSvc, db: db}, nil
}

func (a *App) Close() error {
	return a.db.Close()
}
