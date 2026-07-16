package cli

import (
	"fmt"

	"github.com/dxo1a/obscusync/internal/service"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [profile]",
	Short: "Scan game profile",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		s := service.New()

		result, err := s.Scan(args[0])
		if err != nil {
			return err
		}

		fmt.Printf(
			"Profile '%s' scanned successfully.\nFound %d files.\n",
			args[0],
			result.FileCount,
		)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
