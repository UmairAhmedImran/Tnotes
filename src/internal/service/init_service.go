package service

import (
	"time"
	"tnotes/internal/model"
)

var (
	ConfigFile string = ".notes/config.json"
	NotesDir   string = ".notes"
	BucketName string = "tnotes"

	settings model.SettingsStruct = model.SettingsStruct{
		Encryprtion: false,
		AutoSave:    true,
		Maxnotes:    100,
	}

	markdown model.MarkdownStruct = model.MarkdownStruct{
		Editor:          "nano",
		SyntaxhighLight: true,
	}

	basicConfData model.BasicConfStruct = model.BasicConfStruct{
		Version:       "1.0.0",
		CreatedAt:     time.Now(),
		Scope:         "project",
		DbFile:        "notes.db",
		DefaultBucket: "tnotes",
		Markdown:      markdown,
		Settings:      settings,
	}
)
