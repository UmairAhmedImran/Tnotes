/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"UmairAhmedImran/internal/service"

	"github.com/spf13/cobra"
)

var recursively bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the notes",
	Long: `List the notes in the current project/directory or list all the
the notes globally (if specified)`,
	Run: func(cmd *cobra.Command, args []string) {
		service.CheckInit()

		service.ListCommand(recursively)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(
		&recursively, "recursively", "r", false, "To print the content inside the title or not.",
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
