package handler

import (
	"go-nat-project/models"
	"go-nat-project/utils"

	"github.com/gofiber/fiber/v2"
)

func GetGlobalProcessedScore(c *fiber.Ctx) error {

	globalScore := []models.GlobalScore{

		{
			Title:   "MATH",
			TitleTh: "คณิตศาสตร์",
			Average: 30.5,
		},
		{
			Title:   "SCI",
			TitleTh: "วิทยาศาสตร์",
			Average: 44.5,
		},
		{
			Title:   "ENG",
			TitleTh: "ภาษาอังกฤษ",
			Average: 55.5,
		},
	}

	return utils.SendSuccess(c, globalScore)

}
