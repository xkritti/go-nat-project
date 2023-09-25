package main

import (
	"log"

	"go-nat-project/database"
	_ "go-nat-project/docs"
	router "go-nat-project/rounter"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault) // default
	router.SetupRoutes(app)

	app.Use(
		func(c *fiber.Ctx) error {
			return c.SendStatus(404) // => 404 "Not Found"
		})

	log.Fatal(app.Listen(":" + "4000"))
}
