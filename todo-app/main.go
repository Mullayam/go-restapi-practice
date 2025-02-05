package main

import (
	"log"

	"github.com/Mullayam/todo-app/config"
	"github.com/Mullayam/todo-app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize DB connection
	db := config.NewConnection()
	if db == nil {
		log.Fatalf("Failed to connect to the database")
	}
	defer db.Close()

	// Create repository instance
	repo := &config.Repository{
		DB: db,
	}

	// Create Fiber app
	app := fiber.New()

	// Initialize routes
	routes.InitRoutes(app, repo)

	// Start the server
	log.Println("Server is running on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
