package services

import (
	"web_app/app/models"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ListUsers returns all users.
// It always returns a non-nil slice, even when empty.
func ListUsers() ([]models.User, error) {
	users := make([]models.User, 0)
	err := mgm.Coll(&models.User{}).SimpleFind(&users, bson.M{})
	return users, err
}

// GetUserByID returns a single user by ObjectID.
func GetUserByID(id primitive.ObjectID) (*models.User, error) {
	user := &models.User{}
	if err := mgm.Coll(user).FindByID(id, user); err != nil {
		return nil, err
	}
	return user, nil
}

