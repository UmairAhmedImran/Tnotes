/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"UmairAhmedImran/internal/tui"
	"UmairAhmedImran/internal/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var notesTitle string

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the notes",
	Long: `View the notes on the current project/directory or view
them globally in the termnal.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("view called")
		if notesTitle == "" {

			model, err := tui.BaseScreen()

			if err != nil {
				fmt.Println("Could not initialized Bubble tea model", err)
				os.Exit(1)
			}

			p := tea.NewProgram(model, tea.WithAltScreen()) // , tui.Model
			if _, err := p.Run(); err != nil {
				fmt.Printf("Error running TUI: %v\n", err)
				os.Exit(1)
			}
		} else {
			utils.ViewCommand(notesTitle)
		}
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	viewCmd.Flags().StringVarP(&notesTitle, "NotesTitle", "k", "", "Note title to get")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
