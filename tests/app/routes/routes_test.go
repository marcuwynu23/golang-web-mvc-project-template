package routes_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"web_app/app/database"
	"web_app/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type dummyRenderer struct{}

func (d *dummyRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	_, err := w.Write([]byte("ok"))
	return err
}

func setupEcho() *echo.Echo {
	e := echo.New()
	e.Renderer = &dummyRenderer{}
	// Initialize database (mgm config). This does not require a running Mongo instance.
	database.Init()
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

	// May be 200 (success) or 500 (DB error); just ensure route exists.
	assert.NotEqual(t, http.StatusNotFound, rec.Code)
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

	body := `{"name":"John Doe","email":"john@example.com","age":30}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Depending on DB availability this may be 200 or 500; just assert it's not 404.
	assert.NotEqual(t, http.StatusNotFound, rec.Code)
}

