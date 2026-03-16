package database

import (
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() {
	dbName := getenvDefault("MONGO_DB_NAME", "ginApp")
	uri := getenvDefault("MONGO_URI", "mongodb://localhost:27017")

	_ = mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(uri))
}

func getenvDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

