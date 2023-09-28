package main

import (
	"go-nat-project/config"
	"go-nat-project/database"
	_ "go-nat-project/docs"
	router "go-nat-project/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title Swagger Example APIS
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

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

	port := config.Config().Port
	log.Fatal(app.Listen(":" + port))
}
