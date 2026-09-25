package platform

import (
	"net/url"
	"strings"
)

// Platform identifies a video source site.
type Platform string

const (
	YouTube   Platform = "youtube"
	Instagram Platform = "instagram"
	TikTok    Platform = "tiktok"
	Twitter   Platform = "twitter"
	Facebook  Platform = "facebook"
	Reddit    Platform = "reddit"
	Vimeo     Platform = "vimeo"
	Bilibili  Platform = "bilibili"
	Generic   Platform = "generic"
)

// Detect identifies the platform from a URL string by matching the hostname.
func Detect(rawURL string) Platform {
	u, err := url.Parse(rawURL)
	if err != nil {
		return Generic
	}
	host := strings.ToLower(u.Hostname())
	switch host {
	case "youtube.com", "www.youtube.com", "youtu.be", "m.youtube.com":
		return YouTube
	case "instagram.com", "www.instagram.com":
		return Instagram
	case "tiktok.com", "www.tiktok.com", "vm.tiktok.com":
		return TikTok
	case "twitter.com", "x.com", "mobile.twitter.com":
		return Twitter
	case "facebook.com", "www.facebook.com", "m.facebook.com", "fb.watch":
		return Facebook
	case "reddit.com", "www.reddit.com", "old.reddit.com", "v.redd.it":
		return Reddit
	case "vimeo.com", "player.vimeo.com":
		return Vimeo
	case "bilibili.com", "www.bilibili.com", "b23.tv":
		return Bilibili
	default:
		return Generic
	}
}

// String returns the platform name.
func (p Platform) String() string { return string(p) }

// NeedsCookies reports whether the platform may require browser cookies.
// This is advisory: public content may work without cookies, but private or
// age-gated content will need them. The caller should hint the user rather
// than auto-inject.
func (p Platform) NeedsCookies() bool {
	switch p {
	case Instagram, Facebook:
		return true
	default:
		return false
	}
}

// ExtraArgs returns platform-specific yt-dlp arguments.
// cookieBrowser is the browser name for --cookies-from-browser (empty = skip).
func (p Platform) ExtraArgs(cookieBrowser string) []string {
	var args []string
	if cookieBrowser != "" {
		args = append(args, "--cookies-from-browser", cookieBrowser)
	}
	return args
}

// Label returns a human-readable display name for the platform.
func (p Platform) Label() string {
	switch p {
	case YouTube:
		return "YouTube"
	case Instagram:
		return "Instagram"
	case TikTok:
		return "TikTok"
	case Twitter:
		return "Twitter"
	case Facebook:
		return "Facebook"
	case Reddit:
		return "Reddit"
	case Vimeo:
		return "Vimeo"
	case Bilibili:
		return "Bilibili"
	default:
		return "Generic"
	}
}
