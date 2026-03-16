package database_test

import (
	"context"
	"testing"

	"web_app/app/database"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TestInit verifies Init can be called and (if MongoDB is available)
// that mgm gets a usable default configuration.
func TestInit(t *testing.T) {
	// Try to connect to a local MongoDB; if not reachable, skip.
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		t.Skipf("skipping Init DB test; cannot connect to MongoDB: %v", err)
		return
	}
	defer client.Disconnect(context.Background())

	// Call Init, which configures mgm.
	database.Init()

	// Verify mgm's default config is non-nil.
	cfg := mgm.DefaultConfig()
	if cfg == nil {
		t.Fatalf("mgm default config is nil after Init")
	}
}

