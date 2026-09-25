package cmd

import (
	"fmt"

	"github.com/bizshuk/vdl/pkg/config"
	"github.com/spf13/cobra"
)

// ConfigCmd prints the application config directory path.
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Show config folder path",
	RunE: func(c *cobra.Command, args []string) error {
		dir, err := config.EnsureDir()
		if err != nil {
			return err
		}
		fmt.Println(dir)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(ConfigCmd)
}
