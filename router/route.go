package router

import (
	user_handler "go-nat-project/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	user_api := api.Group("/user")
	user_api.Get("/auth", user_handler.GetUser)
	// user_api.Get("/get_user", user_handler.GetUser)
	user_api.Post("/upload_user", user_handler.UploadUserExcel)
	// auth_api := api.Group("/auth")
	// score_api := api.Group("/score")
	// upload_api := api.Group("/upload")
}
