package main

import (
	"github.com/codeinceo/maur-trivia/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Let code and build!")
	// })

	app.Listen(":3000")
}
