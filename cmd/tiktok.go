package cmd

import (
	"github.com/spf13/cobra"
)

// TiktokCmd represents the first-layer command group for TikTok.
var TiktokCmd = &cobra.Command{
	Use:   "tiktok [flags] <url>...",
	Short: "Download videos from TikTok",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			return c.Help()
		}
		return RunDownload(args)
	},
}

func init() {
	RootCmd.AddCommand(TiktokCmd)
}
