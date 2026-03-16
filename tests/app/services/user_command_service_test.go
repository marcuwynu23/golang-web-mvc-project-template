package services_test

import (
	"context"
	"testing"

	"web_app/app/models"
	"web_app/app/services"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initCommandTestDB(t *testing.T) {
	t.Helper()

	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		t.Skipf("skipping service command DB test; cannot connect to MongoDB: %v", err)
		return
	}

	t.Cleanup(func() {
		_ = client.Disconnect(context.Background())
	})

	if err := mgm.SetDefaultConfig(nil, "serviceCommandDB", clientOpts); err != nil {
		t.Fatalf("failed to set mgm default config: %v", err)
	}
}

func TestCreateAndUpdateAndDeleteUser(t *testing.T) {
	initCommandTestDB(t)

	// Create
	u := &models.User{
		Name:  "Initial",
		Email: "initial@example.com",
		Age:   25,
	}
	if err := services.CreateUser(u); err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	if u.ID.IsZero() {
		t.Fatalf("expected user ID to be set after CreateUser")
	}

	// Update
	payload := &models.User{
		Name: "Updated",
		Age:  30,
	}
	updated, err := services.UpdateUser(u.ID.(primitive.ObjectID), payload)
	if err != nil {
		t.Fatalf("UpdateUser failed: %v", err)
	}
	if updated.Name != "Updated" || updated.Age != 30 {
		t.Fatalf("expected updated fields, got %+v", updated)
	}

	// Delete
	if err := services.DeleteUser(u.ID.(primitive.ObjectID)); err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}
}

