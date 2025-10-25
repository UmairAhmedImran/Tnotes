/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"UmairAhmedImran/internal/utils"
	"UmairAhmedImran/internal/tui/addmodel"

	"time"
    "fmt"
    "os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	title   string
	content string
	dbData  utils.BoltDbStruct

	notes utils.NotesStruct = utils.NotesStruct{
		Id:        uuid.NewString(),
		CreatedAt: time.Now(),
		Content:   content,
		Format:    "markdown",
	}

	BoltStruct utils.BoltDbStruct = utils.BoltDbStruct{
		Title:     title,
		CreatedAt: time.Now(),
		Notes:     []utils.NotesStruct{notes},
	}
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new notes",
	Long: `Add new notes according to your need such as
you can add notes in the current project/directory or you can add golbally`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CheckInit()

		titleProvided := false
		contentProvided := false

    	titleFlag := cmd.Flags().Lookup("title")

		if titleFlag != nil && titleFlag.Changed {
			titleProvided = true
    	}

		contentFlag := cmd.Flags().Lookup("content")

		if contentFlag != nil && contentFlag.Changed {
			fmt.Println("Content is provided by user")
			contentProvided = true
		}

		if titleProvided && !contentProvided {
			currentModel := addmodel.New()
			p := tea.NewProgram(currentModel)
			m, err := p.Run()

			if err != nil {
				fmt.Printf("Error running TUI: %v\n", err)
				os.Exit(1)
			}

			if currentModel, ok := m.(addmodel.Model); ok && currentModel.Value() != "" {
				fmt.Println("Current Model Value:", currentModel.Value())
				BoltStruct.Notes[len(BoltStruct.Notes) - 1].Content = currentModel.Value()
			}


		}
		utils.AddCommand(title, BoltStruct)
	},
}

func init() {

	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(
		&title, "title", "t", "", "Note title",
	)
	addCmd.Flags().StringVarP(
		&BoltStruct.Notes[len(BoltStruct.Notes)-1].Content,
		"content", "c", "", "Note content",
	)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
