package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c *CLI) newScanCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "scan [profile]",
		Short: "Scan game profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(
			cmd *cobra.Command,
			args []string,
		) error {
			result, err := c.service.Scan(args[0])
			if err != nil {
				return err
			}

			fmt.Printf(
				"Profile '%s' scanned successfully.\nFound %d files.\nManifest: %s\n",
				args[0],
				result.FileCount,
				result.Manifest,
			)

			return nil
		},
	}
}
