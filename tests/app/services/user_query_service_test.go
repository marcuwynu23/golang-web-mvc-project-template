package services_test

import (
	"context"
	"testing"

	"web_app/app/services"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// initTestDB sets up a temporary Mongo database for service tests, or skips if not available.
func initTestDB(t *testing.T) {
	t.Helper()

	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		t.Skipf("skipping service DB test; cannot connect to MongoDB: %v", err)
		return
	}

	t.Cleanup(func() {
		_ = client.Disconnect(context.Background())
	})

	if err := mgm.SetDefaultConfig(nil, "serviceTestDB", clientOpts); err != nil {
		t.Fatalf("failed to set mgm default config: %v", err)
	}
}

func TestListUsersEmpty(t *testing.T) {
	initTestDB(t)

	users, err := services.ListUsers()
	if err != nil {
		t.Fatalf("ListUsers returned error: %v", err)
	}
	if users == nil {
		t.Fatalf("expected non-nil slice")
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	initTestDB(t)

	randomID := primitive.NewObjectID()
	user, err := services.GetUserByID(randomID)
	if err == nil {
		t.Fatalf("expected error for non-existent user, got user: %+v", user)
	}
}

