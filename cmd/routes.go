package main

import (
	"github.com/codeinceo/maur-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/:id", handlers.GetFact)
	app.Post("/", handlers.CreateFact)
	app.Put("/:id", handlers.UpdateFact)
	app.Delete("/:id", handlers.DeleteFact)
}
