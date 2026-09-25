package generic

import (
	"context"
	"io"

	"github.com/bizshuk/video-downloader/pkg/download"
	"github.com/bizshuk/video-downloader/svc"
)

func init() {
	svc.Register(&Service{})
}

// Service implements svc.Platform for generic fallback URLs.
type Service struct{}

func (s *Service) Name() string  { return "generic" }
func (s *Service) Label() string { return "Generic" }

func (s *Service) Match(rawURL string) bool {
	return rawURL != ""
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
