package database_test

import (
	"testing"

	"web_app/app/database"
)

// TestInit simply verifies that Init can be called without panicking.
// It does not require a running MongoDB instance.
func TestInit(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Init panicked: %v", r)
		}
	}()

	database.Init()
}

