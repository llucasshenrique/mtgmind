/*
Copyright © 2022 Lucas Henrique <llucasshenrique@gmail.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "MTG Mind database actions",
	Long:  `Run the prepare to create the local database or clear to remove it.`,
}

func init() {
	rootCmd.AddCommand(databaseCmd)
}
