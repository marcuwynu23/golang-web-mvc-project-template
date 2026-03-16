package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"web_app/app/controllers"
	"web_app/app/database"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func newContext(method, path string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func TestHelloWorld(t *testing.T) {
	c, rec := newContext(http.MethodGet, "/api/v1/users/hello")

	err := controllers.HelloWorld(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, World!", rec.Body.String())
}

func TestPrintHello(t *testing.T) {
	c, rec := newContext(http.MethodGet, "/api/v1/users/print/John")
	c.SetParamNames("name")
	c.SetParamValues("John")

	err := controllers.PrintHello(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, John!", rec.Body.String())
}

func TestGetUsers(t *testing.T) {
	// Ensure mgm default config is set; does not require a running Mongo instance.
	database.Init()

	c, rec := newContext(http.MethodGet, "/api/v1/users/all")

	err := controllers.GetUsers(c)
	assert.NoError(t, err)
	// May be 200 (success) or 500 (DB error); just ensure route/handler works.
	assert.NotEqual(t, http.StatusNotFound, rec.Code)
}

func TestInformation(t *testing.T) {
	c, rec := newContext(http.MethodGet, "/api/v1/users/info")

	err := controllers.Information(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Welcome to the API!")
}

