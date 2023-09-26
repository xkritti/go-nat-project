package user_handler

import (
	"errors"
	"fmt"
	"go-nat-project/database"
	"go-nat-project/models"
	user_models "go-nat-project/models"
	"go-nat-project/utils"

	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

// can get all user and filter by year and level_type and major_type

func GetAllUser(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []user_models.Users
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

	result := user_models.User{}
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
	file, err := c.FormFile("file")
	db := database.DB.Db

	if err != nil {
		return errors.New("Uploading Failed!")
	}

	src, err := file.Open()

	if err != nil {
		return errors.New("Source Invalid")
	}

	defer src.Close()

	dest, err := os.Create(file.Filename)

	defer dest.Close()

	_, err = io.Copy(dest, src)

	if err != nil {
		return err
	}

	fmt.Println("Upload Excel")

	excelResult, err := excelize.OpenFile(file.Filename)

	if err != nil {
		return err
	}

	sheetName := excelResult.GetSheetName(0)
	rows, err := excelResult.GetRows(sheetName)
	var user user_models.User

	for i := 2; i < len(rows); i++ {

		cell := fmt.Sprintf("F%d", i)
		cid, err := excelResult.GetCellValue(sheetName, cell)

		if err != nil {
			return err
		}

		user.Cid = utils.GetSha256Enc(cid)

		cell = fmt.Sprintf("C%d", i)
		prefix, _ := excelResult.GetCellValue(sheetName, cell)
		user.Prefix = prefix

		cell = fmt.Sprintf("A%d", i)
		name, _ := excelResult.GetCellValue(sheetName, cell)
		user.Name = name

		cell = fmt.Sprintf("H%d", i)
		levelType, _ := excelResult.GetCellValue(sheetName, cell)
		user.LevelType = levelType

		cell = fmt.Sprintf("J%d", i)
		competitionType, _ := excelResult.GetCellValue(sheetName, cell)
		user.CompetitionType = competitionType

		cell = fmt.Sprintf("K%d", i)
		examLocation, _ := excelResult.GetCellValue(sheetName, cell)
		user.ExamLocation = examLocation

		cell = fmt.Sprintf("L%d", i)
		school, _ := excelResult.GetCellValue(sheetName, cell)
		user.School = school

		cell = fmt.Sprintf("M%d", i)
		province, _ := excelResult.GetCellValue(sheetName, cell)
		user.Province = province

		cell = fmt.Sprintf("O%d", i)
		sector, _ := excelResult.GetCellValue(sheetName, cell)
		user.Sector = sector

		cell = fmt.Sprintf("H%d", i)
		level, _ := excelResult.GetCellValue(sheetName, cell)
		user.Level = level

		result := db.Create(&user)

		fmt.Println(result)

	}

	return c.JSON(models.CommonResponse{
		Code: 1000,
		Data: "Upload Complete",
	})
}
