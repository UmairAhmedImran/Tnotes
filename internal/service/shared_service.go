package service

import (
	"UmairAhmedImran/internal/utils"
	"UmairAhmedImran/internal/db"
	"os"
)

func CheckInit() error {
	if _, err := os.Stat(NotesDir); os.IsNotExist(err) {
		utils.ShowingError.Println("Tnotes not Initialized. Please run `tnotes init` first.")
		os.Exit(1)
	}


	if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
		utils.ShowingInfo.Println("Config file does not exists creating basic config file")
		// going to add two options here to re init or to create a new config?
		CreateConfigFile()
	}


	if _, err := os.Stat(db.DbFile); os.IsNotExist(err) {
		utils.ShowingInfo.Println("DB file does not exists creating new DB file")
		// going to add two options here to re init or to create a new db?
		CreateDbFile()
	}

	return nil
}