package services

import (
	"web_app/app/models"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser persists a new user.
func CreateUser(user *models.User) error {
	return mgm.Coll(user).Create(user)
}

// UpdateUser updates an existing user with the given payload.
// Only non-zero / non-empty fields are applied.
func UpdateUser(id primitive.ObjectID, payload *models.User) (*models.User, error) {
	user := &models.User{}
	if err := mgm.Coll(user).FindByID(id, user); err != nil {
		return nil, err
	}

	if payload.Name != "" {
		user.Name = payload.Name
	}
	if payload.Email != "" {
		user.Email = payload.Email
	}
	if payload.Age != 0 {
		user.Age = payload.Age
	}

	if err := mgm.Coll(user).Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by ID.
func DeleteUser(id primitive.ObjectID) error {
	user := &models.User{}
	if err := mgm.Coll(user).FindByID(id, user); err != nil {
		return err
	}
	return mgm.Coll(user).Delete(user)
}

