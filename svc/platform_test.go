package svc_test

import (
	"testing"

	"github.com/bizshuk/vdl/svc"
	_ "github.com/bizshuk/vdl/svc/bilibili"
	_ "github.com/bizshuk/vdl/svc/facebook"
	_ "github.com/bizshuk/vdl/svc/generic"
	_ "github.com/bizshuk/vdl/svc/instagram"
	_ "github.com/bizshuk/vdl/svc/tiktok"
	_ "github.com/bizshuk/vdl/svc/twitter"
	_ "github.com/bizshuk/vdl/svc/youtube"
)

func TestDetect(t *testing.T) {
	tests := []struct {
		url      string
		wantName string
	}{
		{"https://www.youtube.com/watch?v=abc123", "youtube"},
		{"https://youtu.be/abc123", "youtube"},
		{"https://www.instagram.com/reel/abc123/", "instagram"},
		{"https://instagram.com/p/abc123/", "instagram"},
		{"https://www.tiktok.com/@user/video/123", "tiktok"},
		{"https://vm.tiktok.com/abc123/", "tiktok"},
		{"https://twitter.com/user/status/123", "twitter"},
		{"https://x.com/user/status/123", "twitter"},
		{"https://www.facebook.com/watch/?v=123", "facebook"},
		{"https://fb.watch/abc123/", "facebook"},
		{"https://www.bilibili.com/video/BV1xx", "bilibili"},
		{"https://b23.tv/abc123", "bilibili"},
		{"https://example.com/video.mp4", "generic"},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			p := svc.Detect(tt.url)
			if p == nil {
				t.Fatalf("Detect(%q) returned nil", tt.url)
			}
			if p.Name() != tt.wantName {
				t.Errorf("Detect(%q).Name() = %q, want %q", tt.url, p.Name(), tt.wantName)
			}
		})
	}
}

func TestGetAndAll(t *testing.T) {
	all := svc.All()
	if len(all) < 7 {
		t.Errorf("All() returned %d platforms, want at least 7", len(all))
	}

	for _, p := range all {
		got, ok := svc.Get(p.Name())
		if !ok || got == nil {
			t.Errorf("Get(%q) failed", p.Name())
		}
	}
}

func TestNeedsCookies(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"youtube", false},
		{"instagram", true},
		{"tiktok", false},
		{"twitter", false},
		{"facebook", true},
		{"bilibili", false},
		{"generic", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, ok := svc.Get(tt.name)
			if !ok {
				t.Fatalf("Get(%q) not found", tt.name)
			}
			if got := p.NeedsCookies(); got != tt.want {
				t.Errorf("%s.NeedsCookies() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestExtraArgs(t *testing.T) {
	p, ok := svc.Get("instagram")
	if !ok {
		t.Fatal("instagram not found")
	}

	args := p.ExtraArgs("")
	if len(args) != 0 {
		t.Errorf("ExtraArgs(\"\") = %v, want empty", args)
	}

	args = p.ExtraArgs("chrome")
	if len(args) != 2 || args[0] != "--cookies-from-browser" || args[1] != "chrome" {
		t.Errorf("ExtraArgs(\"chrome\") = %v, want [--cookies-from-browser chrome]", args)
	}
}
