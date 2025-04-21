package router

import (
	"github.com/gramkow-dev/split-it/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/todos", handlers.GetTodos)
	app.Get("/todos/:id", handlers.GetTodo)
	app.Patch("/todos/:id", handlers.ToggleTodoStatus)
	app.Post("/todos", handlers.AddTodo)
}
