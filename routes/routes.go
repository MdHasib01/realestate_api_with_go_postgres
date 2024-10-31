package routes

import (
	"github.com/MdHasib01/realestate_api_with_go_postgre/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/todos", controllers.GetTodos)
	api.Post("/todos", controllers.CreateTodo)
}
