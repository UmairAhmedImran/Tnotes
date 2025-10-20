package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

  "github.com/fatih/color"
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

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		fmt.Println("DB file does not exists creating new DB file")
		// going to add two options here to re init or to create a new db?
		createDbFile()
	} else {
		fmt.Println("DB file is present moving forward")
	}

	return nil
}

func AddCommand(title string, dbData BoltDbStruct) error {

	dbData.Title = title
	db, err := bolt.Open(dbFile, 0600, nil)
	fmt.Println("opening the DB")
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
		fmt.Println("Putting in the DB")

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
			fmt.Printf("value from boltdb: %s\n", string(value))
			err := json.Unmarshal(value, &jsonValue)
			if err != nil {
				return err
			}

			jsonValue.Notes = append(jsonValue.Notes, dbData.Notes...)

			fmt.Printf("value after unMarshal: %s\n", jsonValue)
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
