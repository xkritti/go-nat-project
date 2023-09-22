package main

import (
	"log"

	_ "go-nat-project/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":3000"))
}
