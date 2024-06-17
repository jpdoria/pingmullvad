package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// Mock the command line arguments.
	os.Args = []string{"cmd", "-type=bridge"}

	// Call the main function.
	main()

	// Check if the serverType flag has been set correctly.
	if *serverType != "bridge" {
		t.Errorf("Expected serverType to be 'bridge', got %v", *serverType)
	}
}
