/*
Copyright © 2022 Lucas Henrique <llucasshenrique@gmail.com>

*/
package cmd

import (
	"log"

	"github.com/llucasshenrique/mtgmind/search"
	"github.com/spf13/cobra"
)

var name string
var databasePath string
var set string

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for cards",
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case set != "":
			log.Println("Searching for cards in set:", set)
			search.GetCardBySet(set, databasePath)
		case name != "":
			log.Println("Searching for card:", name)
			search.GetCardByName(name, databasePath)
		default:
			log.Println("Searching for card:", args[0])
			search.GetCardByName(args[0], databasePath)
		}

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&name, "name", "n", "", "name to search")
	searchCmd.Flags().StringVarP(&databasePath, "database", "d", "cards.sqlite", "database to search")
	searchCmd.Flags().StringVarP(&set, "set", "s", "", "set to search")
}
