package handler

import (
	"errors"
	"fmt"
	"go-nat-project/database"
	"go-nat-project/models"
	"go-nat-project/utils"

	"github.com/gofiber/fiber/v2"
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

	result := models.User{}
	err = db.First(&result, "cid = ?", utils.GetSha256Enc(payload.Cid)).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.SendCommonError(c, models.CommonError{
				Code: 2001,
				ErrorData: models.ApiError{
					ErrorTitle:   "Not Found",
					ErrorMessage: "User Not Found",
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
	excelResult, sheetName, rows, err := utils.ExcelReader(filename, 0)

	if err != nil {
		return err
	}

	db := database.DB.Db
	var user models.User
	rows = 10
	for i := 2; i < rows; i++ {
		cid, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("F%d", i))
		fmt.Println(cid)
		prefix, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("C%d", i))
		name, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("A%d", i))
		levelType, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("H%d", i))
		competitionType, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("J%d", i))
		examLocation, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("K%d", i))
		school, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("L%d", i))
		province, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("M%d", i))
		sector, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("O%d", i))
		level, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("H%d", i))

		user.Cid = utils.GetSha256Enc(cid)
		user.Prefix = prefix
		user.Name = name
		user.LevelType = levelType
		user.CompetitionType = competitionType
		user.ExamLocation = examLocation
		user.School = school
		user.Province = province
		user.Sector = sector
		user.Level = level
		result := db.Create(&user)

		fmt.Println(result)

	}

	utils.DeleteFile(filename)

	return c.JSON(models.CommonResponse{
		Code: 1000,
		Data: "Upload Complete",
	})
}
