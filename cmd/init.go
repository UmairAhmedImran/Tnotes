/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
  "log"

  "github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize the notes",
	Long: `init is used to initialize the notes in the project or it
can be initialized globally.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
    
    // get current working Directory
		dir, err := os.Getwd()

		if err != nil {
			fmt.Println("Error getting current Directory:", err)
			return
		}
    
    // check for the .notes folder if not exist initialize one
		notesDir := ".notes"
    configPath := ".notes/config.json"
		if _, err := os.Stat(notesDir); os.IsNotExist(err) {
			os.Mkdir(notesDir, os.ModePerm)

      _, err := os.Create(configPath)
      if err != nil {
        log.Fatalf("Error creating the config file")
      }

			fmt.Println("Notes Directory created at:", dir)
		} else {
			fmt.Println("Notes Directory already exists.")
		}
    
    // create the db file 
    db, err := bolt.Open(notesDir + "/notes.db", 0600, nil) 
    
    if err != nil {
      log.Fatal(err)
    }
    defer db.Close()

    // creating a read write transactions 
    
    db.Update(func(tx *bolt.Tx) error {
    _, err := tx.CreateBucketIfNotExists([]byte("MyNotes"))
    if err != nil {
      return fmt.Errorf("Create bucket: %s", err)
    }
    return nil 
    })
  },
}



func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
