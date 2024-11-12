package models

import (
	"context"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
	Age              int    `json:"age" bson:"age"`
}

// CreateIndex creates a unique index on the Email field
func (u *User) CreateIndex() error {
	collection := mgm.Coll(u)
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},              // Index by Email
		Options: options.Index().SetUnique(true), // Set unique
	}

	// Create the index if it doesn't exist
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	return err
}

// BeforeSave is a hook that runs before saving the User
func (u *User) BeforeSave() error {
	err := u.CreateIndex()
	if err != nil {
		log.Println("Error creating index:", err)
		return err
	}
	return nil
}
