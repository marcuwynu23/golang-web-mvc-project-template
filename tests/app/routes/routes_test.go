package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"web_app/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupEcho() *echo.Echo {
	e := echo.New()
	routes.RoutesRegister(e)
	return e
}

func TestRoutesRootRedirect(t *testing.T) {
	e := setupEcho()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
}

func TestPageHome(t *testing.T) {
	e := setupEcho()

	req := httptest.NewRequest(http.MethodGet, "/page/home", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestApiV1RootRedirect(t *testing.T) {
	e := setupEcho()

	req := httptest.NewRequest(http.MethodGet, "/api/v1", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
}

func TestUsersRedirect(t *testing.T) {
	e := setupEcho()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
}

func TestUsersInfo(t *testing.T) {
	e := setupEcho()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/info", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUsersAll(t *testing.T) {
	e := setupEcho()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/all", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUsersHello(t *testing.T) {
	e := setupEcho()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/hello", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUsersCreate(t *testing.T) {
	e := setupEcho()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/create", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Depending on DB availability this may be 200 or 500; just assert it's not 404.
	assert.NotEqual(t, http.StatusNotFound, rec.Code)
}

