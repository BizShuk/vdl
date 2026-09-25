package cmd

import (
	"github.com/spf13/cobra"
)

// FacebookCmd represents the first-layer command group for Facebook.
var FacebookCmd = &cobra.Command{
	Use:     "facebook [flags] <url>...",
	Aliases: []string{"fb"},
	Short:   "Download videos from Facebook",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			return c.Help()
		}
		return RunDownload(args)
	},
}

func init() {
	RootCmd.AddCommand(FacebookCmd)
}
