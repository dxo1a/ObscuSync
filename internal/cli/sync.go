package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c *CLI) newSyncCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "sync [profile]",
		Short: "Sync local files with remote server manifest",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := c.service.Sync(args[0])
			if err != nil {
				return err
			}

			fmt.Printf(
				"Sync complete.\nDownloaded: %d\nDeleted: %d\nUnchanged: %d\n",
				result.Downloaded,
				result.Deleted,
				result.Unchanged,
			)

			return nil
		},
	}
}
