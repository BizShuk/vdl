package cmd

import (
	"fmt"

	"github.com/bizshuk/video-downloader/pkg/config"
	"github.com/bizshuk/video-downloader/pkg/tui"
	"github.com/spf13/cobra"
)

// MonitorCmd launches the interactive download browser TUI.
var MonitorCmd = &cobra.Command{
	Use:     "monitor",
	Aliases: []string{"m"},
	Short:   "Interactive TUI to browse downloaded files",
	RunE: func(c *cobra.Command, args []string) error {
		dir, err := config.EnsureDataDir()
		if err != nil {
			return fmt.Errorf("ensure data directory: %w", err)
		}
		return tui.Run(dir)
	},
}

func init() {
	RootCmd.AddCommand(MonitorCmd)
}
