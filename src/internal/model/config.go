package model

import (
	"time"
)

type SettingsStruct struct {
	Encryprtion bool `json:"encryprtion"`
	AutoSave    bool `json:"auto_save"`
	Maxnotes    int  `json:"max_notes"`
}

type MarkdownStruct struct {
	Editor          string `json:"editor"`
	SyntaxhighLight bool   `json:"syntax_high_light"`
}

type BasicConfStruct struct {
	Version       string         `json:"name"`
	CreatedAt     time.Time      `json:"created_at"`
	Scope         string         `json:"scope"`
	DbFile        string         `json:"db_file"`
	DefaultBucket string         `json:"notes.db"`
	Markdown      MarkdownStruct `json:"markdown"`
	Settings      SettingsStruct `json:"settings"`
}
