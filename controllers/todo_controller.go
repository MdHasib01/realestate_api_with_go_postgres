package controllers

import (
	"context"

	"github.com/MdHasib01/realestate_api_with_go_postgre/database"
	"github.com/MdHasib01/realestate_api_with_go_postgre/models"
	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	// Query to get all todos
	rows, err := database.DB.Query(context.Background(), "SELECT id, title, status FROM todos")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch todos"})
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Status); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error scanning todo"})
		}
		todos = append(todos, todo)
	}
	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	_, err := database.DB.Exec(context.Background(), "INSERT INTO todos (title, status) VALUES ($1, $2)", todo.Title, todo.Status)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to create todo"})
	}

	return c.Status(201).JSON(todo)
}
