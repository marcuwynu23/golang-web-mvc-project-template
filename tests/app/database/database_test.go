package database_test

import (
	"testing"

	"web_app/app/database"
)

// This test currently just verifies that Init can be called without panicking.
// You can extend it later with a test MongoDB instance.
func TestInit(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Init panicked: %v", r)
		}
	}()
	database.Init()
}

