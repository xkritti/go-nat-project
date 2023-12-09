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

	upload_api := api.Group("/upload-data")
	upload_api.Post("/user-info", handler.UploadUserExcel)
	upload_api.Post("/user-info/update", handler.UpdateUserExcel)

	upload_api.Post("/score", handler.UploadScore)
	upload_api.Post("/score/update", handler.UpdateScore)

	analytic_api := api.Group("/analytic")
	analytic_api.Post("/get-math-iaar", handler.GetMathAnalytic)
	analytic_api.Post("/get-sci-iaar", handler.GetSciAnalytic)
	analytic_api.Post("/get-eng-iaar", handler.GetEngAnalytic)

	globalScoreApi := api.Group("/global-score")
	globalScoreApi.Get("", handler.GetGlobalProcessedScore)

	upload_api.Post("/stat/avg-by-subject", handler.UploadAvgScoreBySubject)
	upload_api.Post("/stat/number-of-competitor-by-province", handler.UploadNumberOfCompetitorByProvince)
	upload_api.Post("/stat/number-of-competitor-by-region", handler.UploadNumberOfCompetitorByRegion)

	score_api := api.Group("/score")
	score_api.Get("", handler.GetScore)

}
