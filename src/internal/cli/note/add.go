/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package note

import (
	"log/slog"
	"tnotes/internal/app"
	"tnotes/internal/note"

	// "tnotes/internal/service"
	// "tnotes/internal/utils"

	"github.com/spf13/cobra"
)

var (
	title   string
	content string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new notes",
	Long: `Add new notes according to your need such as
you can add notes in the current project/directory or you can add golbally`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("[NOTE][CLI]: Function called")

		app, err := app.GetApp()

		if err != nil {
			slog.Error("[NOTE][CLI]: Got some errr", "error", err)
			return
		}

		titleProvided := false
		contentProvided := false

		titleFlag := cmd.Flags().Lookup("title")

		if titleFlag != nil && titleFlag.Changed {
			titleProvided = true
		}

		contentFlag := cmd.Flags().Lookup("content")

		if contentFlag != nil && contentFlag.Changed {
			slog.Debug("[NOTE][CLI]: Content is provided by user")
			// utils.ShowingInfo.Println("Content is provided by user")
			contentProvided = true
		}

		slog.Debug("[NOTE][CLI]: ", "title", title)
		slog.Debug("[NOTE][CLI]: ", "content", content)
		slog.Debug("[NOTE][CLI]: ", "titleFlag", titleFlag)
		slog.Debug("[NOTE][CLI]: ", "contentFlag", contentFlag)

		if titleProvided && contentProvided {
			slog.Debug("Condition passess")
			err := app.Notes.AddNote(cmd.Context(), note.CreateNote{Title: title, Content: content})

			if err != nil {
				slog.Error("[NOTE][CLI]: ", "error", err)
			}
		}
	},
}

func init() {

	addCmd.Flags().StringVarP(
		&title, "title", "t", "", "Note title",
	)
	addCmd.Flags().StringVarP(
		&content,
		"content", "c", "", "Note content",
	)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
