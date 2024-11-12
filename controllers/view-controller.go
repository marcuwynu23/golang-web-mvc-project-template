package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func ShowPage(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"title": "Welcome",
		"msg":   "Hello, Echo!",
	})
}
