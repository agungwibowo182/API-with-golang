package main

import (
	"tugas-bootcamp/config"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New()

	// Load environment variables from .env file
	// Note: You may need to use a library like godotenv to load the .env file
	// Reference: https://pkg.go.dev/github.com/joho/godotenv
	// Example: godotenv.Load()

	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Setup routes
	config.SetupRoutes(app, db)
}
