package middleware

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Middleware is a struct to hold middleware functions
type TemplateRegistry struct {
	templates *template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// LoadTemplates loads all HTML templates from the specified directory and its subdirectories
func LoadTemplates(dir string) (*template.Template, error) {
	// Check if the directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, nil // Return nil if the directory does not exist
	}

	var files []string

	// Walk through the directory and collect all .html files
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Parse all collected template files
	return template.ParseFiles(files...)
}

// SetMiddleware sets up middleware for the Echo instance
func SetMiddleware(e *echo.Echo) (*os.File, error) {
	// Set up logging to a file
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err // Return the error instead of logging it here
	}

	// Create a multi-writer to log to both file and stdout
	multiWriter := io.MultiWriter(file, os.Stdout)
	e.Logger.SetOutput(multiWriter)

	// Use the default logger middleware
	e.Use(middleware.Logger())

	// Set up CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"*"}, // Allow all origins; you can specify specific origins here
		AllowMethods:  []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:  []string{echo.HeaderContentType, echo.HeaderAccept},
		ExposeHeaders: []string{"X-My-Custom-Header"},
		MaxAge:        86400, // Optional: how long the results of a preflight request can be cached
	}))
	// view template middleware
	templates, err := LoadTemplates("views")
	if err != nil {
		e.Logger.Fatal("Failed to load templates:", err)
	}
	if templates != nil {
		e.Renderer = &TemplateRegistry{
			templates: templates,
		}
	} else {
		e.Logger.Warn("No templates loaded: 'views' directory does not exist.")
	}

	return file, nil // Return the file handle for later use
}
