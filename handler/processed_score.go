package handler

import (
	"go-nat-project/database"
	"go-nat-project/models"
	"go-nat-project/utils"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type getAvgScoreRequest struct {
	Year       string `json:"year" validate:"required"`
	LevelRange string `json:"level-range" validate:"required"`
}

type getAvgScoreResponse struct {
	Year       string                       `json:"year"`
	LevelRange string                       `json:"level_range"`
	Scores     map[string]*avgScoreResponse `json:"scores"`
}

type avgScoreResponse struct {
	AvgScore float64 `json:"avg_score"`
	MaxScore float64 `json:"max_score"`
	MinScore float64 `json:"min_score"`
}

func GetAvgScoreByLevelRange(c *fiber.Ctx) error {

	req := &getAvgScoreRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return utils.SendCommonError(c, models.CommonError{
			Code: 4000,
			ErrorData: models.ApiError{
				ErrorTitle:   "Failed to parse request",
				ErrorMessage: err.Error(),
			},
		})
	}

	err = validator.New().Struct(req)
	if err != nil {
		return utils.SendCommonError(c, models.CommonError{
			Code: 4000,
			ErrorData: models.ApiError{
				ErrorTitle:   "Invalid Request",
				ErrorMessage: err.Error(),
			},
		})
	}

	db := database.DB.Db

	avgScore := []models.AvgScoreBySubject{}
	err = db.Where("year = ? AND level_range = ?", req.Year, req.LevelRange).Find(&avgScore).Error
	if err != nil {
		return utils.SendCommonError(c, models.CommonError{
			Code: 5000,
			ErrorData: models.ApiError{
				ErrorTitle:   "Failed to query database",
				ErrorMessage: err.Error(),
			},
		})
	}

	avgScores := map[string]*avgScoreResponse{}
	for _, v := range avgScore {
		avgScores[strings.ToLower(v.Subject)] = &avgScoreResponse{
			AvgScore: v.AvgScore,
			MaxScore: v.MaxScore,
			MinScore: v.MinScore,
		}
	}

	resp := &getAvgScoreResponse{
		Year:       req.Year,
		LevelRange: req.LevelRange,
		Scores:     avgScores,
	}

	return utils.SendSuccess(c, resp)
}

type getNumberOfCompByProvinceRequest struct {
	Year       string `json:"year" validate:"required"`
	LevelRange string `json:"level_range" validate:"required"`
	Province   string `json:"province" validate:"required"`
}

type getNumberOfCompByProvinceResponse struct {
	Year       string           `json:"year" `
	LevelRange string           `json:"level_range"`
	Province   string           `json:"province" `
	Scores     map[string]int64 `json:"scores"`
}

func GetNumberOfCompByProvince(c *fiber.Ctx) error {

	req := &getNumberOfCompByProvinceRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return utils.SendCommonError(c, models.CommonError{
			Code: 4000,
			ErrorData: models.ApiError{
				ErrorTitle:   "Failed to parse request",
				ErrorMessage: err.Error(),
			},
		})
	}

	err = validator.New().Struct(req)
	if err != nil {
		return utils.SendCommonError(c, models.CommonError{
			Code: 4000,
			ErrorData: models.ApiError{
				ErrorTitle:   "Invalid Request",
				ErrorMessage: err.Error(),
			},
		})
	}

	db := database.DB.Db

	numOfCompList := []models.NumberOfCompetitorByProvince{}

	err = db.Where("year = ? AND level_range = ? AND province = ?", req.Year, req.LevelRange, req.Province).Find(&numOfCompList).Error
	if err != nil {
		return utils.SendCommonError(c, models.CommonError{
			Code: 5000,
			ErrorData: models.ApiError{
				ErrorTitle:   "Failed to query database",
				ErrorMessage: err.Error(),
			},
		})
	}

	scores := map[string]int64{}
	for _, v := range numOfCompList {
		scores[strings.ToLower(v.Subject)] = v.NumberOfCompetitor
	}

	resp := &getNumberOfCompByProvinceResponse{
		Year:       req.Year,
		LevelRange: req.LevelRange,
		Province:   req.Province,
		Scores:     scores,
	}

	return utils.SendSuccess(c, resp)
}

type getNumberOfCompByRegionRequest struct {
	Year       string `json:"year" validate:"required"`
	LevelRange string `json:"level_range" validate:"required"`
	Region     string `json:"region" validate:"required"`
}

type getNumberOfCompByRegionResponse struct {
	Year       string           `json:"year" validate:"required"`
	LevelRange string           `json:"level_range" validate:"required"`
	Region     string           `json:"region" validate:"required"`
	Scores     map[string]int64 `json:"scores"`
}

func GetNumberOfCompByRegion(c *fiber.Ctx) error {

	req := &getNumberOfCompByRegionRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return utils.SendCommonError(c, models.CommonError{
			Code: 4000,
			ErrorData: models.ApiError{
				ErrorTitle:   "Failed to parse request",
				ErrorMessage: err.Error(),
			},
		})
	}

	err = validator.New().Struct(req)
	if err != nil {
		return utils.SendCommonError(c, models.CommonError{
			Code: 4000,
			ErrorData: models.ApiError{
				ErrorTitle:   "Invalid Request",
				ErrorMessage: err.Error(),
			},
		})
	}

	db := database.DB.Db

	numOfCompList := []models.NumberOfCompetitorByRegion{}

	err = db.Where("year = ? AND level_range = ? AND region = ?", req.Year, req.LevelRange, req.Region).Find(&numOfCompList).Error
	if err != nil {
		return utils.SendCommonError(c, models.CommonError{
			Code: 5000,
			ErrorData: models.ApiError{
				ErrorTitle:   "Failed to query database",
				ErrorMessage: err.Error(),
			},
		})
	}

	scores := map[string]int64{}
	for _, v := range numOfCompList {
		scores[strings.ToLower(v.Subject)] = v.NumberOfCompetitor
	}

	resp := &getNumberOfCompByRegionResponse{
		Year:       req.Year,
		LevelRange: req.LevelRange,
		Region:     req.Region,
		Scores:     scores,
	}

	return utils.SendSuccess(c, resp)
}
