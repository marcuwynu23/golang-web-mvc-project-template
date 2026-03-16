package models_test

import (
	"testing"

	"web_app/app/models"
)

// This test only checks that User type compiles and can be instantiated.
func TestUserModelInstantiation(t *testing.T) {
	_ = &models.User{
		Name:  "Test",
		Email: "test@example.com",
		Age:   30,
	}
}

