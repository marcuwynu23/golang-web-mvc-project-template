package controllers

import (
  "net/http"
  "github.com/labstack/echo/v4"
  
)
// redirect to /info
func RedirectToInfo(c echo.Context) error {
	return c.Redirect(http.StatusTemporaryRedirect, "/v1/users/info")
}

// print hello world
func PrintHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, "+c.Param("name")+"!")
}

// function to print hello world
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// function to get users
func GetUsers(c echo.Context) error {
	users := map[string]interface{}{
		"users": []map[string]string{
			{"id": "1", "name": "John Doe"},
			{"id": "2", "name": "Jane Doe"},
			{"id": "3", "name": "Mark Doe"},
		},
	}
	return c.JSON(http.StatusOK, users)
}

// function to print information about the API
func Information(c echo.Context) error {
	message := map[string]interface{}{
		"message": "Welcome to the API!",
		"version": "1.0.0",
		"authors": []string{"John Doe", "Jane Doe"},
		"license": "MIT",
	}
	return c.JSON(http.StatusOK, message)
}


