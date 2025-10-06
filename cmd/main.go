package main

import (
	"app/adapters/primary/middlewares"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

const ServiceName = "Go Hexagonal Architecture"

func main() {
	app := fiber.New()

	// Add CORS middleware with default configuration
	app.Use(adaptor.HTTPMiddleware(middlewares.CorsMiddleware(nil)))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to " + ServiceName)
	})

	log.Fatal(app.Listen(":3000"))
}
