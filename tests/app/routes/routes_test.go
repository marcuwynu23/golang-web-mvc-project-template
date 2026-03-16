package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"web_app/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRoutesRegisterRoot(t *testing.T) {
	e := echo.New()
	routes.RoutesRegister(e)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := e.Router().Find(http.MethodGet, "/", c)
	if assert.NotNil(t, h) {
		err := h(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
	}
}

