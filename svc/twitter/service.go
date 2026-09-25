package twitter

import (
	"context"
	"io"
	"net/url"
	"strings"

	"github.com/bizshuk/video-downloader/pkg/download"
	"github.com/bizshuk/video-downloader/svc"
)

func init() {
	svc.Register(&Service{})
}

// Service implements svc.Platform for Twitter/X.
type Service struct{}

func (s *Service) Name() string  { return "twitter" }
func (s *Service) Label() string { return "Twitter/X" }

func (s *Service) Match(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil || u.Host == "" {
		return false
	}
	host := strings.ToLower(u.Hostname())
	return host == "twitter.com" || strings.HasSuffix(host, ".twitter.com") ||
		host == "x.com" || strings.HasSuffix(host, ".x.com")
}

func (s *Service) NeedsCookies() bool           { return false }
func (s *Service) DefaultCookieBrowser() string { return "" }

func (s *Service) ExtraArgs(cookieBrowser string) []string {
	var args []string
	if cookieBrowser != "" {
		args = append(args, "--cookies-from-browser", cookieBrowser)
	}
	return args
}

func (s *Service) Download(ctx context.Context, opts svc.Options, stdout, stderr io.Writer) error {
	req := download.Request{
		URL:       opts.URL,
		Type:      opts.Type,
		Quality:   opts.Quality,
		OutputDir: opts.OutputDir,
		ExtraArgs: s.ExtraArgs(opts.CookieBrowser),
	}
	return download.Run(ctx, req, stdout, stderr)
}
