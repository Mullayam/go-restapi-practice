package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"name"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var todos = []Todo{}

func GetTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}
func CreateTodo(c *fiber.Ctx) error {
	todo := &Todo{}
	if err := c.BodyParser(todo); err != nil {
		c.Status(400).JSON(&fiber.Map{"error": err.Error()})
		return err
	}
	todo.ID = len(todos) + 1
	todos = append(todos, *todo)
	return c.JSON(todo)
}
func UpdateTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(404).SendString("Todo not found")
	}
	todo := &Todo{}
	if err := c.BodyParser(todo); err != nil {
		return err
	}

	for i, t := range todos {
		if t.ID == id {
			todos[i] = *todo
			return c.JSON(todo)
		}
	}
	return c.JSON(todo)

}
func DeleteTodo(c *fiber.Ctx) error {
	return c.JSON(todos)
}
