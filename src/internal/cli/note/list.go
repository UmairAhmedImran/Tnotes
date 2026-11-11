/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package note

import (
	"log/slog"
	"tnotes/internal/app"

	"github.com/spf13/cobra"
)

var detailed bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the notes",
	Long: `List the notes in the current project/directory or list all the
the notes globally (if specified)`,
	Run: func(cmd *cobra.Command, args []string) {
		// service.CheckInit()

		app, err := app.GetApp()

		if err != nil {
			slog.Error("[NOTE][CLI] could not get app: ", "error", err)
			return
		}

		app.Notes.ListNote(cmd.Context(), detailed)
	},
}

func init() {

	listCmd.Flags().BoolVarP(
		&detailed, "detailed", "d", false, "To print the content inside the title or not.",
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
