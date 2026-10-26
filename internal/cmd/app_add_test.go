package cmd

import "testing"

func TestAddLauncherAppRejectsUnsafeID(t *testing.T) {
	if err := addLauncherApp("../x"); err == nil {
		t.Fatal("expected error")
	}
}

func TestAddLauncherAppMissingBinary(t *testing.T) {
	if err := addLauncherApp("agent-terminal-not-a-real-binary"); err == nil {
		t.Fatal("expected error")
	}
}
