package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/fatih/color"
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

var (
	notesDir   string = ".notes"
	configFile string = ".notes/config.json"
	dbFile     string = ".notes/notes.db"
	bucketName string = "tnotes"
	//content    string
	//title      string

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
	jsonValue BoltDbStruct = BoltDbStruct{}
    showingWarning = color.RGB(255, 128, 0)
    showingError = color.RGB(255, 255, 0)
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

	fmt.Println("Created config file")
	return nil
}

func createDbFile() (*bolt.DB, error) {
	// create the db file
	filePermission := 0600
	db, err := bolt.Open(dbFile, os.FileMode(filePermission), nil)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Created Database File")

	// creating a read write transactions

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tnotes"))
		if err != nil {
			return fmt.Errorf("Create bucket: %s", err)
		}
		return nil
	})

	fmt.Println("DB Bucket created successfully")

	return db, nil
}

func CallInit() error {
	_, err := os.Getwd()

	if err != nil {
		log.Fatal("Error getting current directory")
	}

	if _, err := os.Stat(notesDir); os.IsNotExist(err) {
		os.Mkdir(notesDir, os.ModePerm)

		createConfigFile()

	} else {
		showingWarning.Println("Notes Directory already exists.")
	}

	createDbFile()

	fmt.Println("Tnotes Initialization Complete")

	return nil
}

func CheckInit() error {
	if _, err := os.Stat(notesDir); os.IsNotExist(err) {
		showingError.Println("Tnotes not Initialized. Please run `tnotes init` first.")
		os.Exit(1)
	} 
	

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("Config file does not exists creating basic config file")
		// going to add two options here to re init or to create a new config?
		createConfigFile()
	} 


	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		fmt.Println("DB file does not exists creating new DB file")
		// going to add two options here to re init or to create a new db?
		createDbFile()
	} 

	return nil
}

func AddCommand(title string, dbData BoltDbStruct) error {

	dbData.Title = title

	filePermission := 0600
	db, err := bolt.Open(dbFile, os.FileMode(filePermission), nil)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))

		if b == nil {
			// ask user to create a new bucket or go for health check or init again?
			return fmt.Errorf("bucket %s does not exists", bucketName)
		}

		value := b.Get([]byte(title))

		if value == nil {
			bytesData, err := json.Marshal(dbData)
			if err != nil {
				return err
			}
			fmt.Printf("Marshal data: %s", bytesData)

			if err := b.Put([]byte(title), bytesData); err != nil {
				return err
			}
		} else {
			err := json.Unmarshal(value, &jsonValue)
			if err != nil {
				return err
			}

			jsonValue.Notes = append(jsonValue.Notes, dbData.Notes...)

			data, err := json.Marshal(jsonValue)
			if err != nil {
				return err
			}
			err = b.Put([]byte(title), data)
			return err
		}
		return nil
	})

}

func ViewCommand(title ...string) error {

  CheckInit()

	if len(title) == 0 {
		db, err := bolt.Open(dbFile, 0600, nil)
		if err != nil {
			return err
		}
		defer db.Close()

		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucketName))
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Printf("key=%s, value=%s\n", k, v)
			}
			return nil
		})

	} else if len(title) == 1 {
		db, err := bolt.Open(dbFile, 0600, nil)
		if err != nil {
			return err
		}
		defer db.Close()

		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucketName))
			value := b.Get([]byte(title[0]))
			if value == nil {
				showingWarning.Println("There is no such title")
			}
			return nil
		})
	} else {
		fmt.Println("Error too many arguments")
	}
	return nil
}

func ListCommand(recursively bool) error {

	db, err := bolt.Open(dbFile, os.FileMode(0600), nil)

	if err != nil {
		fmt.Println("Error opening the DB File")
	}

	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		cursor := bucket.Cursor()
		fmt.Println()
		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			fmt.Printf("Title: %s\n", key)

			if recursively {
				var curRecord BoltDbStruct
				err := json.Unmarshal(value, &curRecord)

				if err != nil {
					fmt.Println("Error loading data from db")
				}

				for noteIndex := 0; noteIndex < len(curRecord.Notes); noteIndex += 1 {

					currentNote := curRecord.Notes[noteIndex]
					noteFirstLine := currentNote.Content
					formattedID := currentNote.Id[:8]
					maxPreviewLength := 40

					if len(noteFirstLine) > maxPreviewLength {
						noteFirstLine = noteFirstLine[:maxPreviewLength]
						noteFirstLine = strings.Join([]string{noteFirstLine, "..."}, "")
					}

					fmt.Printf("   ↳ ID: %s | Content: %s \n", formattedID, noteFirstLine)
				}

				fmt.Println(strings.Repeat("-", 120))

			}

			fmt.Println()
		}

		return nil
	})

	return nil
}

