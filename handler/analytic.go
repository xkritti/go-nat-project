package handler

import (
	"fmt"
	"go-nat-project/models"
	"go-nat-project/utils"

	"github.com/gofiber/fiber/v2"
)

func getClassification(prefix string, score float64) string {
	if score >= 0 && score <= 25.0 {
		return fmt.Sprintf("%s1", prefix)
	}
	if score >= 25.01 && score <= 49.99 {
		return fmt.Sprintf("%s2", prefix)
	}
	if score >= 50.00 && score <= 69.99 {
		return fmt.Sprintf("%s3", prefix)
	}
	if score >= 70.00 && score <= 84.99 {
		return fmt.Sprintf("%s4", prefix)
	}
	if score >= 85.00 && score <= 100.0 {
		return fmt.Sprintf("%s5", prefix)
	}

	return "-"
}
func GetMathAnalytic(c *fiber.Ctx) error {

	payload := &models.GetMathAnalyticRequest{}
	err := c.BodyParser(payload)
	if err != nil {
		return c.SendStatus(400)
	}

	result := &models.MathAnalytic{}
	result.Classification = getClassification("M", float64(payload.ScorePercentage))
	// CAL Part

	if payload.CalPartScore >= 0 && payload.CalPartScore <= 5.65 {
		result.Parts.Calculation = "M1"
	}
	if payload.CalPartScore >= 5.66 && payload.CalPartScore <= 11.28 {
		result.Parts.Calculation = "M2"
	}
	if payload.CalPartScore >= 11.29 && payload.CalPartScore <= 15.80 {
		result.Parts.Calculation = "M3"
	}
	if payload.CalPartScore >= 15.81 && payload.CalPartScore <= 19.19 {
		result.Parts.Calculation = "M4"
	}
	if payload.CalPartScore >= 19.20 && payload.CalPartScore <= 22.60 {
		result.Parts.Calculation = "M5"
	}

	// Problem Solving

	if payload.ProblemPartScore >= 0 && payload.ProblemPartScore <= 13.16 {
		result.Parts.ProblemSolution = "M1"
	}
	if payload.ProblemPartScore >= 13.17 && payload.ProblemPartScore <= 26.27 {
		result.Parts.ProblemSolution = "M2"
	}
	if payload.ProblemPartScore >= 26.28 && payload.ProblemPartScore <= 36.80 {
		result.Parts.ProblemSolution = "M3"
	}
	if payload.ProblemPartScore >= 36.81 && payload.ProblemPartScore <= 44.70 {
		result.Parts.ProblemSolution = "M4"
	}
	if payload.ProblemPartScore >= 44.71 && payload.ProblemPartScore <= 52.65 {
		result.Parts.ProblemSolution = "M5"
	}

	// Applied Part

	if payload.AppliedPartScore >= 0 && payload.AppliedPartScore <= 6.19 {
		result.Parts.Appliation = "M1"
	}
	if payload.AppliedPartScore >= 6.20 && payload.AppliedPartScore <= 12.35 {
		result.Parts.Appliation = "M2"
	}
	if payload.AppliedPartScore >= 12.36 && payload.AppliedPartScore <= 17.30 {
		result.Parts.Appliation = "M3"
	}
	if payload.AppliedPartScore >= 17.31 && payload.AppliedPartScore <= 21.02 {
		result.Parts.Appliation = "M4"
	}
	if payload.AppliedPartScore >= 21.03 && payload.AppliedPartScore <= 24.75 {
		result.Parts.Appliation = "M5"
	}

	return utils.SendSuccess(c, result)

}

func GetSciAnalytic(c *fiber.Ctx) error {
	payload := &models.GetSciAnalyticRequest{}
	err := c.BodyParser(payload)
	if err != nil {
		return c.SendStatus(400)
	}

	result := &models.SciAnalytic{}

	result.Classification = getClassification("S", float64(payload.ScorePercentage))
	// Lesson Part

	if payload.LessonPartScore >= 0 && payload.LessonPartScore <= 20.13 {
		result.Parts.Lesson = "S1"
	}
	if payload.LessonPartScore >= 20.14 && payload.LessonPartScore <= 40.17 {
		result.Parts.Lesson = "S2"
	}
	if payload.LessonPartScore >= 40.18 && payload.LessonPartScore <= 56.27 {
		result.Parts.Lesson = "S3"
	}
	if payload.LessonPartScore >= 56.28 && payload.LessonPartScore <= 68.34 {
		result.Parts.Lesson = "S4"
	}
	if payload.LessonPartScore >= 68.35 && payload.LessonPartScore <= 80.5 {
		result.Parts.Lesson = "S5"
	}

	// Applied Part

	if payload.AppliedPartScore >= 0 && payload.AppliedPartScore <= 4.88 {
		result.Parts.Appliation = "S1"
	}
	if payload.AppliedPartScore >= 4.89 && payload.AppliedPartScore <= 9.73 {
		result.Parts.Appliation = "S2"
	}
	if payload.AppliedPartScore >= 9.74 && payload.AppliedPartScore <= 13.63 {
		result.Parts.Appliation = "S3"
	}
	if payload.AppliedPartScore >= 13.64 && payload.AppliedPartScore <= 16.56 {
		result.Parts.Appliation = "S4"
	}
	if payload.AppliedPartScore >= 16.57 && payload.AppliedPartScore <= 1950 {
		result.Parts.Appliation = "S5"
	}

	return utils.SendSuccess(c, result)
}

func GetEngAnalytic(c *fiber.Ctx) error {
	payload := &models.GetEngAnalyticRequest{}
	err := c.BodyParser(payload)
	if err != nil {
		return c.SendStatus(400)
	}

	result := &models.EngAnalytic{}

	result.Classification = getClassification("E", float64(payload.ScorePercentage))
	return nil

}
