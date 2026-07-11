package cli

import (
	"fmt"

	"github.com/dxo1a/obscusync/internal/manifest"
	"github.com/dxo1a/obscusync/internal/scanner"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [directory]",
	Short: "Scan directory (choose main game folder)",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		s := scanner.New()

		files, err := s.Scan(args[0])
		if err != nil {
			return err
		}

		builder := manifest.New()
		m := builder.Build(files)
		if err := manifest.Save("manifest.json", m); err != nil {
			return err
		}

		fmt.Println()

		fmt.Printf("Found %d files\n\n", len(files))
		fmt.Println("Manifest saved to manifest.json")

		// for _, file := range files {
		// 	fmt.Printf(
		// 		"%s\nSize: %d\nSHA256: %s\n\n",
		// 		file.Path,
		// 		file.Size,
		// 		file.SHA256,
		// 	)
		// }
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
