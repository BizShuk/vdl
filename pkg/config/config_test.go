package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDirRespectsXDG(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", "/tmp/xdg-test")
	got, err := Dir()
	if err != nil {
		t.Fatalf("Dir() error: %v", err)
	}
	want := filepath.Join("/tmp/xdg-test", "vdl")
	if got != want {
		t.Errorf("Dir() = %q, want %q", got, want)
	}
}

func TestDirFallsBackToHome(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", "")
	got, err := Dir()
	if err != nil {
		t.Fatalf("Dir() error: %v", err)
	}
	home, _ := os.UserHomeDir()
	want := filepath.Join(home, ".config", "vdl")
	if got != want {
		t.Errorf("Dir() = %q, want %q", got, want)
	}
}

func TestDataDirIsSubpath(t *testing.T) {
	got, err := DataDir()
	if err != nil {
		t.Fatalf("DataDir() error: %v", err)
	}
	if !strings.HasSuffix(got, filepath.Join("vdl", "data")) {
		t.Errorf("DataDir() = %q, should end with vdl/data", got)
	}
}

func TestEnsureDataDirCreatesDirectory(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", tmp)

	dir, err := EnsureDataDir()
	if err != nil {
		t.Fatalf("EnsureDataDir() error: %v", err)
	}
	info, err := os.Stat(dir)
	if err != nil {
		t.Fatalf("stat(%q) error: %v", dir, err)
	}
	if !info.IsDir() {
		t.Errorf("%q is not a directory", dir)
	}
}
