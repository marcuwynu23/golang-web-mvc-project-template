package controllers

import (
  "net/http"
  "github.com/labstack/echo/v4"
  
)

// function to handle errors
func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	// Send the JSON response without returning an error
	c.JSON(code, map[string]string{
		"error": err.Error(),
	})
}
