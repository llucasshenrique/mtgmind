/*
Copyright © 2022 Lucas Henrique <llucasshenrique@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mtgmind",
	Short: "MTG Mind is an Magic The Gathering card database and deck builder",
	Long: `MTG Mind is an Magic The Gathering card database and deck builder.

It is a tool to help you build decks and decks for Magic The Gathering`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
