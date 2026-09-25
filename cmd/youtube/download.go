package youtube

import (
	"github.com/bizshuk/vdl/cmd"
	"github.com/spf13/cobra"
)

// DownloadCmd is the second-layer download verb for YouTube.
var DownloadCmd = &cobra.Command{
	Use:   "download [flags] <url>...",
	Short: "Download specified YouTube video URLs",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			return c.Help()
		}
		return cmd.RunDownload(args)
	},
}

func init() {
	cmd.YoutubeCmd.AddCommand(DownloadCmd)
}
