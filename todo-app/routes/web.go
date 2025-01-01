package routes

import (
	"github.com/Mullayam/todo-app/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	api.Get("/todos", controllers.GetTodos)
	api.Post("/todos", controllers.CreateTodo)
	api.Patch("/todos/:id", controllers.UpdateTodo)
}
