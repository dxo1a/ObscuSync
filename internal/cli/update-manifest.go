package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c *CLI) newUpdateManifestCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "update-manifest [profile]",
		Short: "Update game manifest",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := c.service.Scan(args[0])
			if err != nil {
				return err
			}

			fmt.Printf(
				"Profile '%s' updated successfully.\nFound %d files.\nManifest: %s\n",
				args[0],
				result.FileCount,
				result.Manifest,
			)

			return nil
		},
	}
}
