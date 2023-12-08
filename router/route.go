package router

import (
	"go-nat-project/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	user_api := api.Group("/user")
	user_api.Post("/auth", handler.GetUser)
	// user_api.Get("/get_user", user_handler.GetUser)
	// auth_api := api.Group("/auth")
	// score_api := api.Group("/score")

	upload_api := api.Group("/upload-data")
	upload_api.Post("/user-info", handler.UploadUserExcel)
	upload_api.Post("/user-info/update", handler.UpdateUserExcel)

	upload_api.Post("/score", handler.UploadScore)
	upload_api.Post("/score/update", handler.UpdateScore)
}
