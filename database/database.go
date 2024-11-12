package database

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() {
	_ = mgm.SetDefaultConfig(nil, "ginApp", options.Client().ApplyURI("mongodb://localhost:27017"))
	
}
