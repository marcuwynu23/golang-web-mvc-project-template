package main

import (
	"log"
	"github.com/labstack/echo/v4"
	"web_app/middleware"
	"web_app/routes"
)

func main() {
	e := echo.New()
	// Set up middleware
	logFile, err := middleware.SetMiddleware(e)
	if err != nil {
		log.Fatalf("error setting up middleware: %v", err)
	}
	defer logFile.Close() // Close the log file when the application exits

	// Register routes
	routes.RoutesRegister(e)
	// Start the server on 0.0.0.0:8080
	if err := e.Start("0.0.0.0:8080"); err != nil {
		e.Logger.Fatal(err)
	}
}