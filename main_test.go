package main

import (
	"testing"

	"github.com/bizshuk/vdl/cmd"
)

func TestMainCommandStructure(t *testing.T) {
	if cmd.RootCmd == nil {
		t.Fatal("RootCmd should not be nil")
	}
	if cmd.RootCmd.Use == "" {
		t.Fatal("RootCmd.Use should not be empty")
	}
}
