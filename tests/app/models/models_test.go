package models_test

import (
	"context"
	"testing"

	"web_app/app/models"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TestUserModelInstantiation checks that User type compiles and can be instantiated.
func TestUserModelInstantiation(t *testing.T) {
	_ = &models.User{
		Name:  "Test",
		Email: "test@example.com",
		Age:   30,
	}
}

// TestUserCreateIndex ensures CreateIndex builds an index model against a test collection.
// It uses a local, in-memory-like Mongo instance if available; otherwise it skips on connection error.
func TestUserCreateIndex(t *testing.T) {
	// Try to connect to a local MongoDB; skip if not available.
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		t.Skipf("skipping DB index test; cannot connect to MongoDB: %v", err)
		return
	}
	defer client.Disconnect(context.Background())

	err = mgm.SetDefaultConfig(nil, "testdb", clientOpts)
	if err != nil {
		t.Fatalf("failed to set mgm default config: %v", err)
	}

	u := &models.User{}
	if err := u.CreateIndex(); err != nil {
		t.Fatalf("CreateIndex failed: %v", err)
	}

	// Verify index exists (by name pattern)
	coll := mgm.Coll(u)
	cursor, err := coll.Indexes().List(context.Background())
	if err != nil {
		t.Fatalf("failed to list indexes: %v", err)
	}
	defer cursor.Close(context.Background())

	found := false
	for cursor.Next(context.Background()) {
		var idx bson.M
		if err := cursor.Decode(&idx); err != nil {
			t.Fatalf("failed to decode index: %v", err)
		}
		if key, ok := idx["key"].(bson.M); ok {
			if _, ok := key["email"]; ok {
				found = true
				break
			}
		}
	}

	if !found {
		t.Fatalf("email index not found on users collection")
	}
}


