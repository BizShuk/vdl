package download

import (
	"strings"
	"testing"
)

func TestRequestValidate(t *testing.T) {
	tests := []struct {
		name    string
		req     Request
		wantErr bool
	}{
		{
			name: "valid mp3",
			req: Request{
				URL: "https://example.com/video", Type: TYPE_MP3,
				Quality: 3, OutputDir: "/tmp",
			},
		},
		{
			name: "valid mp4",
			req: Request{
				URL: "https://example.com/video", Type: TYPE_MP4,
				Quality: 5, OutputDir: "/tmp",
			},
		},
		{
			name:    "missing URL",
			req:     Request{URL: "", Type: TYPE_MP3, Quality: 3},
			wantErr: true,
		},
		{
			name:    "bad type",
			req:     Request{URL: "https://x.com/v", Type: "avi", Quality: 3},
			wantErr: true,
		},
		{
			name:    "bad quality",
			req:     Request{URL: "https://x.com/v", Type: TYPE_MP3, Quality: 99},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRequestArgs(t *testing.T) {
	req := Request{
		URL:       "https://instagram.com/reel/abc",
		Type:      TYPE_MP4,
		Quality:   3,
		OutputDir: "/out",
		ExtraArgs: []string{"--cookies-from-browser", "chrome"},
	}
	args := req.Args("/tmp/path")
	joined := strings.Join(args, " ")
	if !strings.Contains(joined, "--no-playlist") {
		t.Error("args missing --no-playlist")
	}
	if !strings.Contains(joined, "https://instagram.com/reel/abc") {
		t.Error("args missing URL")
	}
	if !strings.Contains(joined, "--format") {
		t.Error("args missing --format for mp4")
	}
	if !strings.Contains(joined, "--cookies-from-browser chrome") {
		t.Error("args missing extra args")
	}
}

func TestRequestArgsURLIsLast(t *testing.T) {
	req := Request{
		URL:       "https://example.com/video",
		Type:      TYPE_MP3,
		Quality:   1,
		OutputDir: "/out",
	}
	args := req.Args("/tmp/path")
	if len(args) == 0 {
		t.Fatal("args is empty")
	}
	if args[len(args)-1] != req.URL {
		t.Errorf("URL should be last arg, got %q", args[len(args)-1])
	}
}
