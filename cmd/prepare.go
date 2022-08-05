/*
Copyright © 2022 Lucas Henrique <llucasshenrique@gmail.com>

*/
package cmd

import (
	"github.com/llucasshenrique/mtgmind/database"
	"github.com/spf13/cobra"
)

// prepareCmd represents the prepare command
var source string
var dbFile string
var force bool
var prepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Prepare the database for usage",
	Long: `Prepare the database for usage.

This command will create the database file if it doesn't exist.
Is needed to download the All Cards json from: https://scryfall.com/docs/api/bulk-data
To upgrade the database use the --force flag
`,
	Run: func(cmd *cobra.Command, args []string) {
		database.PrepareDatabase(source, dbFile, force)
	},
}

func init() {
	databaseCmd.AddCommand(prepareCmd)
	prepareCmd.Flags().StringVarP(&source, "file", "f", "", "Path to the file")
	prepareCmd.Flags().StringVarP(&dbFile, "db", "d", "cards.sqlite", "Path to the database file")
	prepareCmd.Flags().BoolVarP(&force, "force", "", false, "Force overwrite of the database")
}
