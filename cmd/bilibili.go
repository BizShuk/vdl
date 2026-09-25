package cmd

import (
	"github.com/spf13/cobra"
)

// BilibiliCmd represents the first-layer command group for Bilibili.
var BilibiliCmd = &cobra.Command{
	Use:     "bilibili [flags] <url>...",
	Aliases: []string{"bili"},
	Short:   "Download videos from Bilibili",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			return c.Help()
		}
		return RunDownload(args)
	},
}

func init() {
	RootCmd.AddCommand(BilibiliCmd)
}
