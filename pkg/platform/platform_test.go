package platform

import (
	"reflect"
	"testing"
)

func TestDetect(t *testing.T) {
	tests := []struct {
		url  string
		want Platform
	}{
		{"https://youtube.com/watch?v=123", YouTube},
		{"https://youtu.be/123", YouTube},
		{"https://m.youtube.com/watch?v=123", YouTube},
		{"https://www.youtube.com/watch?v=abc123", YouTube},
		{"https://www.instagram.com/p/123/", Instagram},
		{"https://vm.tiktok.com/123/", TikTok},
		{"https://x.com/user/status/123", Twitter},
		{"https://fb.watch/123/", Facebook},
		{"https://v.redd.it/123", Reddit},
		{"https://vimeo.com/123", Vimeo},
		{"https://b23.tv/123", Bilibili},
		{"https://example.com", Generic},
		{"://invalid", Generic},
	}
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			if got := Detect(tt.url); got != tt.want {
				t.Errorf("Detect(%q) = %v, want %v", tt.url, got, tt.want)
			}
		})
	}
}

func TestNeedsCookies(t *testing.T) {
	tests := []struct {
		p    Platform
		want bool
	}{
		{Instagram, true},
		{TikTok, true},
		{Facebook, true},
		{YouTube, false},
		{Twitter, false},
		{Reddit, false},
		{Vimeo, false},
		{Bilibili, false},
		{Generic, false},
	}
	for _, tt := range tests {
		t.Run(string(tt.p), func(t *testing.T) {
			if got := tt.p.NeedsCookies(); got != tt.want {
				t.Errorf("%v.NeedsCookies() = %v, want %v", tt.p, got, tt.want)
			}
		})
	}
}

func TestExtraArgs(t *testing.T) {
	tests := []struct {
		name          string
		p             Platform
		cookieBrowser string
		want          []string
	}{
		{"instagram without cookies", Instagram, "", []string{"--extractor-args", "instagram:api_type=graphql"}},
		{"instagram with cookies", Instagram, "chrome", []string{"--cookies-from-browser", "chrome", "--extractor-args", "instagram:api_type=graphql"}},
		{"youtube with cookies", YouTube, "firefox", []string{"--cookies-from-browser", "firefox"}},
		{"youtube without cookies", YouTube, "", nil},
		{"tiktok with cookies", TikTok, "safari", []string{"--cookies-from-browser", "safari"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.ExtraArgs(tt.cookieBrowser)
			if len(got) == 0 && len(tt.want) == 0 {
				return // Both nil or empty
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v.ExtraArgs(%q) = %v, want %v", tt.p, tt.cookieBrowser, got, tt.want)
			}
		})
	}
}

func TestLabel(t *testing.T) {
	tests := []struct {
		p    Platform
		want string
	}{
		{YouTube, "YouTube"},
		{Instagram, "Instagram"},
		{TikTok, "TikTok"},
		{Twitter, "Twitter"},
		{Facebook, "Facebook"},
		{Reddit, "Reddit"},
		{Vimeo, "Vimeo"},
		{Bilibili, "Bilibili"},
		{Generic, "Generic"},
	}
	for _, tt := range tests {
		t.Run(string(tt.p), func(t *testing.T) {
			if got := tt.p.Label(); got != tt.want {
				t.Errorf("%v.Label() = %v, want %v", tt.p, got, tt.want)
			}
		})
	}
}
