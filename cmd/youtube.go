package cmd

import (
	"github.com/spf13/cobra"
)

// YoutubeCmd represents the first-layer command group for YouTube.
var YoutubeCmd = &cobra.Command{
	Use:     "youtube [flags] <url>...",
	Aliases: []string{"yt"},
	Short:   "Download videos and audio from YouTube",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			return c.Help()
		}
		return RunDownload(args)
	},
}

func init() {
	RootCmd.AddCommand(YoutubeCmd)
}
