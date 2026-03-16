package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"web_app/app/middleware"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSetMiddleware(t *testing.T) {
	// Ensure no leftover log file interferes
	_ = os.Remove("server.log")

	e := echo.New()
	logFile, err := middleware.SetMiddleware(e)
	assert.NoError(t, err)
	assert.NotNil(t, logFile)

	// Logger middleware should set a non-nil logger output
	assert.NotNil(t, e.Logger)

	// CORS middleware should allow a simple OPTIONS preflight
	req := httptest.NewRequest(http.MethodOptions, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNoContent, rec.Code)

	_ = logFile.Close()
}

