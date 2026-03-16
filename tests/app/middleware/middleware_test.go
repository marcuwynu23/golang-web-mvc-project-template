package middleware_test

import (
	"testing"

	"web_app/app/middleware"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSetMiddleware(t *testing.T) {
	e := echo.New()
	logFile, err := middleware.SetMiddleware(e)
	assert.NoError(t, err)
	if logFile != nil {
		_ = logFile.Close()
	}
}

