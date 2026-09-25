package instagram

import (
	"github.com/bizshuk/video-downloader/cmd"
	"github.com/spf13/cobra"
)

// DownloadCmd is the second-layer download verb for Instagram.
var DownloadCmd = &cobra.Command{
	Use:   "download [flags] <url>...",
	Short: "Download specified Instagram reel or post URLs",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			return c.Help()
		}
		return cmd.RunDownload(args)
	},
}

func init() {
	cmd.InstagramCmd.AddCommand(DownloadCmd)
}
