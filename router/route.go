package router

import (
	"go-nat-project/handler"
	"go-nat-project/middlaware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	api := app.Group("/api")

	user_api := api.Group("/user")
	user_api.Post("/auth", handler.GetUser)

	analytic_api := api.Group("/analytic")
	analytic_api.Post("/get-math-iaar", handler.GetMathAnalytic)
	analytic_api.Post("/get-sci-iaar", handler.GetSciAnalytic)
	analytic_api.Post("/get-eng-iaar", handler.GetEngAnalytic)
	analytic_api.Post("/get-iaar", handler.GetIaarData)

	statApi := api.Group("/stat")
	statApi.Get("/avg-score-by-level-range", handler.GetAvgScoreByLevelRange)
	statApi.Get("/number-of-comp-by-province", handler.GetNumberOfCompByProvince)
	statApi.Get("/number-of-comp-by-region", handler.GetNumberOfCompByRegion)

	score_api := api.Group("/score")
	score_api.Get("", handler.GetScore)

	upload_api := api.Group("/upload-data", middlaware.IsAdmin)
	upload_api.Post("/user-info", handler.UploadUserExcel)
	upload_api.Post("/user-info/update", handler.UpdateUserExcel)
	upload_api.Post("/score", handler.UploadScore)
	upload_api.Post("/score/update", handler.UpdateScore)
	upload_api.Post("/stat/avg-by-subject", handler.UploadAvgScoreBySubject)
	upload_api.Post("/stat/number-of-competitor-by-province", handler.UploadNumberOfCompetitorByProvince)
	upload_api.Post("/stat/number-of-competitor-by-region", handler.UploadNumberOfCompetitorByRegion)
}
