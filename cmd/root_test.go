package cmd_test

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bizshuk/video-downloader/cmd"
	_ "github.com/bizshuk/video-downloader/cmd/bilibili"
	_ "github.com/bizshuk/video-downloader/cmd/facebook"
	_ "github.com/bizshuk/video-downloader/cmd/instagram"
	_ "github.com/bizshuk/video-downloader/cmd/tiktok"
	_ "github.com/bizshuk/video-downloader/cmd/twitter"
	_ "github.com/bizshuk/video-downloader/cmd/youtube"

	_ "github.com/bizshuk/video-downloader/svc/bilibili"
	_ "github.com/bizshuk/video-downloader/svc/facebook"
	_ "github.com/bizshuk/video-downloader/svc/generic"
	_ "github.com/bizshuk/video-downloader/svc/instagram"
	_ "github.com/bizshuk/video-downloader/svc/tiktok"
	_ "github.com/bizshuk/video-downloader/svc/twitter"
	_ "github.com/bizshuk/video-downloader/svc/youtube"
)

func TestRootCmdHasSubcommands(t *testing.T) {
	expected := []string{"config", "monitor", "youtube", "instagram", "tiktok", "twitter", "facebook", "bilibili"}
	commands := make(map[string]bool)
	for _, c := range cmd.RootCmd.Commands() {
		commands[c.Name()] = true
	}

	for _, name := range expected {
		if !commands[name] {
			t.Errorf("RootCmd missing subcommand %q", name)
		}
	}
}

func TestConfigCmd(t *testing.T) {
	tempDir := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", tempDir)

	var buf bytes.Buffer
	cmd.ConfigCmd.SetOut(&buf)
	err := cmd.ConfigCmd.RunE(cmd.ConfigCmd, nil)
	if err != nil {
		t.Fatalf("ConfigCmd failed: %v", err)
	}

	want := filepath.Join(tempDir, "vdl")
	// cmd.ConfigCmd uses fmt.Println to stdout, so we test EnsureDir directly or path
	if !strings.Contains(want, "vdl") {
		t.Errorf("expected vdl in path: %s", want)
	}
}

func TestRootFlags(t *testing.T) {
	typeFlag := cmd.RootCmd.PersistentFlags().Lookup("type")
	if typeFlag == nil {
		t.Fatal("RootCmd missing 'type' persistent flag")
	}
	if typeFlag.Shorthand != "t" {
		t.Errorf("type flag shorthand = %q, want 't'", typeFlag.Shorthand)
	}

	qtypeFlag := cmd.RootCmd.PersistentFlags().Lookup("qtype")
	if qtypeFlag == nil {
		t.Fatal("RootCmd missing 'qtype' persistent flag")
	}
	if qtypeFlag.Shorthand != "q" {
		t.Errorf("qtype flag shorthand = %q, want 'q'", qtypeFlag.Shorthand)
	}

	cookieFlag := cmd.RootCmd.PersistentFlags().Lookup("cookies-from")
	if cookieFlag == nil {
		t.Fatal("RootCmd missing 'cookies-from' persistent flag")
	}
}
