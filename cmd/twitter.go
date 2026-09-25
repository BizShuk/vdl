package cmd

import (
	"github.com/spf13/cobra"
)

// TwitterCmd represents the first-layer command group for Twitter/X.
var TwitterCmd = &cobra.Command{
	Use:     "twitter [flags] <url>...",
	Aliases: []string{"x"},
	Short:   "Download videos from Twitter/X",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			return c.Help()
		}
		return RunDownload(args)
	},
}

func init() {
	RootCmd.AddCommand(TwitterCmd)
}
