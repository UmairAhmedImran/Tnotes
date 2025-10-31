package service

import (
	"UmairAhmedImran/internal/model"
	"UmairAhmedImran/internal/utils"
	"UmairAhmedImran/internal/db"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
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

func CreateConfigFile() error {
	conFile, err := os.Create(ConfigFile)
	if err != nil {
		log.Fatalf("Error creating the config file")
	}
	defer conFile.Close()

	encoder := json.NewEncoder(conFile)
	err = encoder.Encode(basicConfData)

	if err != nil {
		utils.ShowingWarning.Printf("Error encoding JSON: %v\n", err)
		os.Exit(1)
	}

	utils.ShowingSuccess.Println("Created config file")
	return nil
}

func CreateDbFile() (*bolt.DB, error) {
	// create the db file
	filePermission := 0600
	db, err := bolt.Open(db.DbFile, os.FileMode(filePermission), nil)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	utils.ShowingSuccess.Println("Created Database File")

	// creating a read write transactions

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	utils.ShowingSuccess.Println("DB Bucket created successfully")

	return db, nil
}

func CallInit() error {
	_, err := os.Getwd()

	if err != nil {
		log.Fatal("Error getting current directory")
	}

	if _, err := os.Stat(NotesDir); os.IsNotExist(err) {
		os.Mkdir(NotesDir, os.ModePerm)

		CreateConfigFile()

	} else {
		utils.ShowingWarning.Println("Notes Directory already exists.")
	}

	CreateDbFile()

	utils.ShowingSuccess.Println("Tnotes Initialization Complete")

	return nil
}
