/*
Copyright © 2022 Lucas Henrique <llucasshenrique@gmail.com>

*/
package cmd

import (
	"github.com/llucasshenrique/mtgmind/search"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var name string
var databasePath string
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for cards",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			name = args[0]
		}
		if name != "" {
			search.GetCardByName(name, databasePath)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&name, "name", "n", "", "name to search")
	searchCmd.Flags().StringVarP(&databasePath, "database", "d", "cards.sqlite", "database to search")
}
