package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

type settingsStruct struct {
	Encryprtion bool `json:"encryprtion"`
	AutoSave    bool `json:"auto_save"`
	Maxnotes    int  `json:"max_notes"`
}

type markdownStruct struct {
	Editor          string `json:"editor"`
	SyntaxhighLight bool   `json:"syntax_high_light"`
}

type basicConfStruct struct {
	Version       string         `json:"name"`
	CreatedAt     time.Time      `json:"created_at"`
	Scope         string         `json:"scope"`
	DbFile        string         `json:"db_file"`
	DefaultBucket string         `json:"notes.db"`
	Markdown      markdownStruct `json:"markdown"`
	Settings      settingsStruct `json:"settings"`
}

var (
	notesDir   string = ".notes"
	configFile string = ".notes/config.json"
	dbFile     string = ".notes/notes.db"
	bucketName string = "tnotes"

	settings settingsStruct = settingsStruct{
		Encryprtion: false,
		AutoSave:    true,
		Maxnotes:    100,
	}

	markdown markdownStruct = markdownStruct{
		Editor:          "nano",
		SyntaxhighLight: true,
	}
	basicConfData basicConfStruct = basicConfStruct{
		Version:       "1.0.0",
		CreatedAt:     time.Now(),
		Scope:         "project",
		DbFile:        "notes.db",
		DefaultBucket: "tnotes",
		Markdown:      markdown,
		Settings:      settings,
	}
)

func createConfigFile() error {
	conFile, err := os.Create(configFile)
	if err != nil {
		log.Fatalf("Error creating the config file")
	}
	defer conFile.Close()

	encoder := json.NewEncoder(conFile)
	err = encoder.Encode(basicConfData)

	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Creating config file")
	return nil
}

func createDbFile() (*bolt.DB, error) {
	// create the db file
	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// creating a read write transactions

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tnotes"))
		if err != nil {
			return fmt.Errorf("Create bucket: %s", err)
		}
		return nil
	})
	return db, nil
}

func CallInit() error {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal("Error getting current directory")
	}

	if _, err := os.Stat(notesDir); os.IsNotExist(err) {
		os.Mkdir(notesDir, os.ModePerm)

		createConfigFile()

		fmt.Println("Notes Directory created at:", dir)
	} else {
		fmt.Println("Notes Directory already exists.")
	}

	createDbFile()

	return nil
}

func CheckInit() error {
	if _, err := os.Stat(notesDir); os.IsNotExist(err) {
		fmt.Println("Tnotes not Initialized")
		os.Exit(1)
	} else {
		fmt.Println("Notes is Initialized moving forward")
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("Config file does not exists creating basic config file")
		// going to add two options here to re init or to create a new config?
		createConfigFile()
	} else {
		fmt.Println("Config file is present moving forward")
	}

	if _, err := os.Stat(DBFile); os.IsNotExist(err) {
		fmt.Println("DB file does not exists creating new DB file")
		// going to add two options here to re init or to create a new db?
		createDbFile()
	} else {
		fmt.Println("DB file is present moving forward")
	}

	return nil
}

func AddCommand(title, content string) error {

	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)

		if b == nil {
			// ask user to create a new bucket or go for health check or init again?
			return fmt.Errorf("bucket %s does not exists", bucketName)
		}

		return b.Put([]byte(key), []byte(value))
	})
}
