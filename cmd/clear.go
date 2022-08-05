/*
Copyright © 2022 Lucas Henrique <llucasshenrique@gmail.com>

*/
package cmd

import (
	"github.com/llucasshenrique/mtgmind/database"
	"github.com/spf13/cobra"
)

var path string
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Deletes the database",
	Long: `Deletes the database file.
This is a destructive operation and cannot be undone.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.ClearDatabase(path)
	},
}

func init() {
	databaseCmd.AddCommand(clearCmd)
	clearCmd.Flags().StringVarP(&path, "path", "p", "cards.sqlite", "Relative path to the database file")
}
