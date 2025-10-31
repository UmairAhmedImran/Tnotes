package model

import (
	"time"
)

type NotesStruct struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
	Format    string    `json:"format"`
}

type BoltDbStruct struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	//tags          tagsStruct      `json:"tags"`  tags can be added later
	Notes []NotesStruct `json:"notes"`
}