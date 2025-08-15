/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
  "os"

  "github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new notes",
	Long: `Add new notes according to your need such as 
you can add notes in the current project/directory or you can add golbally`,
  Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		notesDir := ".notes"
    
    if _, err := os.Stat(notesDir); os.IsNotExist(err) {
      fmt.Printf("Initialize the notes first before adding: %s", err)
    }

    //configPath := ".notes/config.json"
    
    _, err := bolt.Open(notesDir, 0600, nil)
    
    if err != nil {
      fmt.Printf("Error opening the db")
    }

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
