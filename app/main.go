package main

import (
	"log"
	"os"

	"web_app/app/database"
	"web_app/app/middleware"
	"web_app/app/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables from .env if present
	if err := godotenv.Load(); err != nil {
		// In production it's fine if .env is missing; only log for visibility.
		log.Printf(".env file not found or could not be loaded: %v", err)
	}

	// Optional: allow overriding listen address via ENV
	addr := os.Getenv("APP_LISTEN_ADDR")
	if addr == "" {
		addr = "0.0.0.0:8080"
	}

	database.Init()

	e := echo.New()
	// Set up middleware
	logFile, err := middleware.SetMiddleware(e)
	if err != nil {
		log.Fatalf("error setting up middleware: %v", err)
	}
	defer logFile.Close() // Close the log file when the application exits

	// Register routes
	routes.RoutesRegister(e)
	// Start the server
	if err := e.Start(addr); err != nil {
		e.Logger.Fatal(err)
	}
}

