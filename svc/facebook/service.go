package facebook

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

// Service implements svc.Platform for Facebook.
type Service struct{}

func (s *Service) Name() string  { return "facebook" }
func (s *Service) Label() string { return "Facebook" }

func (s *Service) Match(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil || u.Host == "" {
		return false
	}
	host := strings.ToLower(u.Hostname())
	return host == "facebook.com" || strings.HasSuffix(host, ".facebook.com") ||
		host == "fb.watch" || strings.HasSuffix(host, ".fb.watch")
}

func (s *Service) NeedsCookies() bool           { return true }
func (s *Service) DefaultCookieBrowser() string { return "chrome" }

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
