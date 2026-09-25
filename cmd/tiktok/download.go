package tiktok

import (
	"github.com/bizshuk/video-downloader/cmd"
	"github.com/spf13/cobra"
)

// DownloadCmd is the second-layer download verb for TikTok.
var DownloadCmd = &cobra.Command{
	Use:   "download [flags] <url>...",
	Short: "Download specified TikTok video URLs",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			return c.Help()
		}
		return cmd.RunDownload(args)
	},
}

func init() {
	cmd.TiktokCmd.AddCommand(DownloadCmd)
}
