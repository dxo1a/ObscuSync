package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gamesync",
	Short: "GameSync synchronizes game mods and configs",
	Long:  "GameSync is a CLI tool for synchronizing game mods and configuration files.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GameSync")
		fmt.Println("Use 'gamesync --help' to see available commands.")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
