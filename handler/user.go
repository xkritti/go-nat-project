package handler

import (
	"errors"
	"fmt"
	"go-nat-project/database"
	"go-nat-project/models"
	"go-nat-project/utils"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// var ErrList []models.Competitor

// can get all user and filter by year and level_type and major_type

func GetAllUser(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []models.Users
	db.Find(&users)
	if len(users) == 0 {
		return c.JSON(models.CommonResponse{
			Code: 2000,
			Data: "OK",
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"data":    users,
			"message": "Get all user success!",
		})
	}
}

type getUserRequest struct {
	Cid string `json:"cid"`
}

func GetUser(c *fiber.Ctx) error {
	db := database.DB.Db

	req := &getUserRequest{}

	err := c.BodyParser(req)
	if err != nil {
		return err
	}

	err = validator.New().Struct(req)
	if err != nil {
		return utils.SendCommonError(c, models.CommonError{
			Code: 2001,
			ErrorData: models.ApiError{
				ErrorTitle:   "ไม่สามารถดำเนินการได้",
				ErrorMessage: "ไม่พบรายชื่อในระบบ",
			},
		})
	}

	result := models.Competitor{}
	hashCid := utils.GetSha256Enc(req.Cid)

	err = db.Where("cid = ?", hashCid).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.SendCommonError(c, models.CommonError{
				Code: 2001,
				ErrorData: models.ApiError{
					ErrorTitle:   "ไม่สามารถดำเนินการได้",
					ErrorMessage: "ไม่พบรายชื่อในระบบ",
				},
			})
		} else {
			return err
		}
	}

	if slr, err := utils.GetShortLevelRange(result.LevelRange); err != nil {
		result.ShortLevelRange = result.LevelRange
	} else {
		result.ShortLevelRange = slr
	}

	return utils.SendSuccess(c, result)
}

func UploadUserExcel(c *fiber.Ctx) error {

	sheetIndex, err := strconv.Atoi(c.FormValue("sheet_index"))
	if err != nil {
		sheetIndex = 0
	}

	filename, err := utils.UploadFileReader(c)
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Upload Failed with %v", err.Error()),
		})
	}
	excelResult, sheetName, _, err := utils.ExcelReader(filename, sheetIndex)
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Upload Failed with %v", err.Error()),
		})
	}

	db := database.DB.Db
	competitorList := []*models.Competitor{}
	rows, err := excelResult.GetRows(sheetName)
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Upload Failed with %v", err.Error()),
		})
	}

	colReader := utils.NewColumnReader(rows[1])
	for i := 3; i <= len(rows); i++ {
		examType, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.ExamType), i))
		name, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Name), i))
		cid, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Cid), i), excelize.Options{RawCellValue: true})
		levelRange, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.LevelRange), i))
		level, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Level), i))
		province, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Province), i))
		region, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Region), i))
		school, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.School), i))
		examLocation, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.ExamLocation), i))

		fmt.Printf("| %s | %s | %s | %s | %s | %s | %s | %s | %s  \n", examType, name, cid, levelRange, level, province, region, school, examLocation)

		competitorList = append(competitorList, &models.Competitor{
			ID:           uuid.New().String(),
			Name:         name,
			Cid:          utils.GetSha256Enc(strings.TrimSpace(cid)),
			ExamType:     examType,
			LevelRange:   levelRange,
			Level:        level,
			Province:     province,
			Region:       region,
			School:       school,
			ExamLocation: examLocation,
		})
	}
	tx := db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(competitorList, 500)
	if tx.Error != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Upload Failed with %v", tx.Error.Error()),
		})
	}

	utils.DeleteFile(filename)

	return c.JSON(models.CommonResponse{
		Code: 1000,
		Data: fmt.Sprintf("Upload Complete total row = %v, insert into db %v record", len(competitorList), tx.RowsAffected),
	})
}

type UpdateUserExcelRequest struct {
	sheetIndex int `json:"sheet_index"`
}

func UpdateUserExcel(c *fiber.Ctx) error {
	filename, err := utils.UploadFileReader(c)
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Update Failed with %v", err.Error()),
		})
	}

	sheetIndex, err := strconv.Atoi(c.FormValue("sheet_index"))
	if err != nil {
		sheetIndex = 0
	}

	excelResult, sheetName, _, err := utils.ExcelReader(filename, sheetIndex)
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Update Failed with %v", err.Error()),
		})
	}

	db := database.DB.Db
	competitorList := []*models.Competitor{}
	rows, err := excelResult.GetRows(sheetName)
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Update Failed with %v", err.Error()),
		})
	}

	colReader := utils.NewColumnReader(rows[1])
	for i := 3; i <= len(rows); i++ {
		examType, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.ExamType), i))
		name, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Name), i))
		cid, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Cid), i), excelize.Options{RawCellValue: true})
		levelRange, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.LevelRange), i))
		level, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Level), i))
		province, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Province), i))
		region, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Region), i))
		school, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.School), i))
		examLocation, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.ExamLocation), i))

		fmt.Printf("| %s | %s | %s | %s | %s | %s | %s | %s | %s  \n", examType, name, cid, levelRange, level, province, region, school, examLocation)

		competitorList = append(competitorList, &models.Competitor{
			Name:         name,
			Cid:          utils.GetSha256Enc(strings.TrimSpace(cid)),
			ExamType:     examType,
			LevelRange:   levelRange,
			Level:        level,
			Province:     province,
			Region:       region,
			School:       school,
			ExamLocation: examLocation,
		})
	}

	errList := []models.Competitor{}
	for _, v := range competitorList {
		updateColumn := []string{}
		if v.ExamType != "" {
			updateColumn = append(updateColumn, "exam_type")
		}
		if v.Name != "" {
			updateColumn = append(updateColumn, "name")
		}
		if v.LevelRange != "" {
			updateColumn = append(updateColumn, "level_range")
		}
		if v.Level != "" {
			updateColumn = append(updateColumn, "level")
		}
		if v.Province != "" {
			updateColumn = append(updateColumn, "province")
		}
		if v.Region != "" {
			updateColumn = append(updateColumn, "region")
		}
		if v.School != "" {
			updateColumn = append(updateColumn, "school")
		}
		if v.ExamLocation != "" {
			updateColumn = append(updateColumn, "exam_location")
		}

		tx := db.Model(&models.Competitor{}).Where("cid = ?", v.Cid).Select(updateColumn).Updates(v)
		if tx.Error != nil {
			errList = append(errList, *v)
		}
		if tx.RowsAffected == 0 || tx.RowsAffected > 1 {
			errList = append(errList, *v)
		}
	}

	utils.DeleteFile(filename)

	return c.JSON(models.CommonResponse{
		Code: 1000,
		Data: struct {
			Message string
			ErrList []models.Competitor
		}{
			Message: fmt.Sprintf("Update Complete total row = %v", len(competitorList)),
			ErrList: errList,
		},
	})
}
