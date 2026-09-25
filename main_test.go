package main

import (
	"bytes"
	"flag"
	"strings"
	"testing"
)

func TestRunConfigWritesPath(t *testing.T) {
	var stdout, stderr bytes.Buffer
	err := runConfig(nil, &stdout, &stderr)
	if err != nil {
		t.Fatalf("runConfig() error: %v", err)
	}
	out := strings.TrimSpace(stdout.String())
	if !strings.Contains(out, "vdl") {
		t.Errorf("runConfig() output = %q, want path containing \"vdl\"", out)
	}
}

func TestRunConfigRejectsExtraArgs(t *testing.T) {
	var stdout, stderr bytes.Buffer
	err := runConfig([]string{"extra"}, &stdout, &stderr)
	if err == nil {
		t.Error("runConfig(extra) = nil, want error")
	}
}

func TestUsageContainsPlatforms(t *testing.T) {
	// Capture usage output by redirecting flag output
	var buf bytes.Buffer
	old := flag.CommandLine.Output()
	flag.CommandLine.SetOutput(&buf)
	defer flag.CommandLine.SetOutput(old)

	usage()
	out := buf.String()
	for _, p := range []string{"YouTube", "Instagram", "TikTok", "Twitter/X", "Facebook", "Reddit", "Vimeo", "Bilibili"} {
		if !strings.Contains(out, p) {
			t.Errorf("usage() missing platform %q", p)
		}
	}
}
