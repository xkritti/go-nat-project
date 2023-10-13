package handler

import (
	"errors"
	"fmt"
	"go-nat-project/database"
	"go-nat-project/models"
	"go-nat-project/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

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

func GetUser(c *fiber.Ctx) error {
	db := database.DB.Db

	payload := struct {
		Cid string `json:"cid"`
	}{}

	err := c.BodyParser(&payload)

	if err != nil {
		return err
	}

	result := models.Competitor{}
	err = db.First(&result, "cid = ?", utils.GetSha256Enc(payload.Cid)).Error

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

	return utils.SendSuccess(c, result)
}

func UploadUserExcel(c *fiber.Ctx) error {
	filename, err := utils.UploadFileReader(c)
	excelResult, sheetName, _, err := utils.ExcelReader(filename, 0)
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: "Upload Failed",
		})
	}

	db := database.DB.Db
	var competitor models.Competitor
	rows, err := excelResult.GetRows("Sheet1")
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: "Upload Failed",
		})
	}
	for i := 2; i < len(rows); i++ {
		examType, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("A%d", i))
		name, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("B%d", i))
		cid, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("C%d", i), excelize.Options{RawCellValue: true})
		fmt.Printf("%d | RAW CID : %s", i, cid)
		levelRange, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("D%d", i))
		level, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("E%d", i))
		province, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("F%d", i))
		region, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("G%d", i))
		school, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("H%d", i))
		examLocation, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("I%d", i))
		competitor.ExamType = examType
		competitor.Name = name
		competitor.Cid = utils.GetSha256Enc(cid)
		competitor.LevelRange = levelRange
		competitor.Level = level
		competitor.Province = province
		competitor.Region = region
		competitor.School = school
		competitor.ExamLocation = examLocation
		fmt.Printf(" | %s | %s | %s | %s | %s | %s | %s | %s | %s  \n", examType, name, cid, levelRange, level, province, region, school, examLocation)
		db.Create(&competitor)

		if err != nil {
			return c.JSON(models.CommonResponse{
				Code: 1001,
				Data: "Upload Failed",
			})

		}

	}

	utils.DeleteFile(filename)

	return c.JSON(models.CommonResponse{
		Code: 1000,
		Data: "Upload Complete",
	})
}
