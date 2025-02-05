package routes

import (
	"github.com/Mullayam/todo-app/config"
	"github.com/Mullayam/todo-app/controllers"
	"github.com/gofiber/fiber/v2"
)

// InitRoutes sets up all application routes
func InitRoutes(app *fiber.App, db *config.Repository) {
	api := app.Group("/api")
	todoController := controllers.NewTodoController(db.DB)
	api.Get("/todos", todoController.GetTodos)
	api.Post("/todos", todoController.CreateTodo)
	api.Patch("/todos/:id", todoController.UpdateTodo)
}
