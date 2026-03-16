package controllers

import (
	"net/http"

	"web_app/app/models"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// redirect to /info
func RedirectToInfo(c echo.Context) error {
	return c.Redirect(http.StatusTemporaryRedirect, "/api/v1/users/info")
}

// print hello world
func PrintHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, "+c.Param("name")+"!")
}

// function to print hello world
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// GetUsers returns all users.
func GetUsers(c echo.Context) error {
	var users []models.User
	err := mgm.Coll(&models.User{}).SimpleFind(&users, bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

// Information returns API meta information.
func Information(c echo.Context) error {
	message := map[string]interface{}{
		"message": "Welcome to the API!",
		"version": "1.0.0",
		"authors": []string{"John Doe", "Jane Doe"},
		"license": "MIT",
	}
	return c.JSON(http.StatusOK, message)
}

// CreateUser creates a new user from JSON payload.
func CreateUser(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	if err := mgm.Coll(user).Create(user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

// GetUser retrieves a single user by ID.
func GetUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	user := &models.User{}
	if err := mgm.Coll(user).FindByID(id, user); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser updates an existing user.
func UpdateUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	user := &models.User{}
	if err := mgm.Coll(user).FindByID(id, user); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
	}

	payload := &models.User{}
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user by ID.
func DeleteUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	user := &models.User{}
	if err := mgm.Coll(user).FindByID(id, user); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
	}

	if err := mgm.Coll(user).Delete(user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

