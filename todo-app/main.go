package main

import (
	"github.com/Mullayam/todo-app/config"
	"github.com/Mullayam/todo-app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	config.NewConnection()

	app := fiber.New()
	routes.InitRoutes(app)

	app.Listen(":3000")
}
