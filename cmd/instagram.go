package cmd

import (
	"github.com/spf13/cobra"
)

// InstagramCmd represents the first-layer command group for Instagram.
var InstagramCmd = &cobra.Command{
	Use:     "instagram [flags] <url>...",
	Aliases: []string{"ig"},
	Short:   "Download reels and posts from Instagram",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			return c.Help()
		}
		return RunDownload(args)
	},
}

func init() {
	RootCmd.AddCommand(InstagramCmd)
}
