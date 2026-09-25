package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bizshuk/video-downloader/pkg/config"
	"github.com/bizshuk/video-downloader/pkg/download"
	"github.com/bizshuk/video-downloader/svc"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	// Root and persistent flags
	MediaTypeFlag string
	QualityFlag   int
	CookieBrowser string
)

// RootCmd is the top-level command.
var RootCmd = &cobra.Command{
	Use:   "vdl [flags] <url>...",
	Short: "vdl: Universal video downloader CLI wrapping yt-dlp",
	Long: `vdl downloads video/audio URLs from YouTube, Instagram, TikTok, Twitter/X,
Facebook, Reddit, Vimeo, Bilibili and more at quality tiers 1~5.
Output contract: stdout carries only the resulting file path, progress goes to stderr.`,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			return c.Help()
		}
		return RunDownload(args)
	},
}

// downloadCmd handles direct URL arguments when dispatched.
var downloadCmd = &cobra.Command{
	Use:    "download [flags] <url>...",
	Short:  "Download specified video URLs",
	Hidden: true,
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			return c.Help()
		}
		return RunDownload(args)
	},
}

// Execute executes the root command, automatically routing URLs to download.
func Execute() {
	if len(os.Args) > 1 {
		// Normalize single-dash flags to double-dash for seamless compatibility with ytdl
		for i := 1; i < len(os.Args); i++ {
			arg := os.Args[i]
			if strings.HasPrefix(arg, "-type") && !strings.HasPrefix(arg, "--type") {
				os.Args[i] = "-" + arg
			} else if strings.HasPrefix(arg, "-qtype") && !strings.HasPrefix(arg, "--qtype") {
				os.Args[i] = "-" + arg
			} else if strings.HasPrefix(arg, "-cookies-from") && !strings.HasPrefix(arg, "--cookies-from") {
				os.Args[i] = "-" + arg
			}
		}

		firstNonFlag := ""
		for _, arg := range os.Args[1:] {
			if !strings.HasPrefix(arg, "-") {
				firstNonFlag = arg
				break
			}
		}
		if firstNonFlag != "" && !IsKnownCommand(firstNonFlag) {
			// Prepend "download" subcommand
			newArgs := make([]string, 0, len(os.Args)+1)
			newArgs = append(newArgs, os.Args[0], "download")
			newArgs = append(newArgs, os.Args[1:]...)
			os.Args = newArgs
		}
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "vdl: %v\n", err)
		os.Exit(1)
	}
}

// IsKnownCommand reports whether name matches a known subcommand or alias.
func IsKnownCommand(name string) bool {
	if name == "help" || name == "completion" {
		return true
	}
	for _, c := range RootCmd.Commands() {
		if c.Name() == name {
			return true
		}
		for _, alias := range c.Aliases {
			if alias == name {
				return true
			}
		}
	}
	return false
}

// RunDownload downloads the given URLs using platform detection and services.
func RunDownload(urls []string) error {
	mediaType, err := download.ParseMediaType(MediaTypeFlag)
	if err != nil {
		return err
	}
	if err := download.ValidateQuality(QualityFlag); err != nil {
		return err
	}

	outputDir, err := config.EnsureDataDir()
	if err != nil {
		return fmt.Errorf("ensure data directory: %w", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	for _, rawURL := range urls {
		p := svc.Detect(rawURL)
		if p == nil {
			return fmt.Errorf("no suitable platform found for URL: %s", rawURL)
		}

		cb := CookieBrowser
		if cb == "none" {
			cb = ""
		}
		if cb == "" && p.NeedsCookies() {
			fmt.Fprintf(os.Stderr, "    hint: %s may need cookies. If download fails, try: --cookies-from chrome\n", p.Label())
		}

		opts := svc.Options{
			URL:           rawURL,
			Type:          mediaType,
			Quality:       QualityFlag,
			OutputDir:     outputDir,
			CookieBrowser: cb,
		}

		fmt.Fprintf(os.Stderr, "==> %s [%s] [%s %s]\n", rawURL, p.Label(), mediaType, download.Describe(mediaType, QualityFlag))
		if err := p.Download(ctx, opts, os.Stdout, os.Stderr); err != nil {
			return err
		}
	}
	return nil
}

func init() {
	RootCmd.SetGlobalNormalizationFunc(normalizeFlags)

	RootCmd.PersistentFlags().StringVarP(&MediaTypeFlag, "type", "t", string(download.TYPE_DEFAULT), "output format: mp3 or mp4")
	RootCmd.PersistentFlags().IntVarP(&QualityFlag, "qtype", "q", download.QUALITY_DEFAULT, "quality tier 1 (lowest) to 5 (highest)")
	RootCmd.PersistentFlags().StringVar(&CookieBrowser, "cookies-from", "", "browser to read cookies from (chrome, firefox, safari, etc.)")

	RootCmd.AddCommand(downloadCmd)
}

// normalizeFlags allows single-dash long flags (-type, -qtype) to match POSIX style flags.
func normalizeFlags(f *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "type":
		return pflag.NormalizedName("type")
	case "qtype":
		return pflag.NormalizedName("qtype")
	case "cookies-from":
		return pflag.NormalizedName("cookies-from")
	}
	return pflag.NormalizedName(name)
}
