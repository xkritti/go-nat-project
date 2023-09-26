package user_handler

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"go-nat-project/database"
	user_models "go-nat-project/models"
	"io"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

// can get all user and filter by year and level_type and major_type
func GetAllUser(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []user_models.Users
	db.Find(&users)
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No user found!",
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
	return c.SendString("Users, World 👋!")
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
		h := sha256.New()

		cell := fmt.Sprintf("F%d", i)
		cid, err := excelResult.GetCellValue(sheetName, cell)

		if err != nil {
			return err
		}

		h.Write([]byte(cid))

		bs := h.Sum(nil)

		hashedCid := fmt.Sprintf("%x", bs)

		user.Cid = hashedCid

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

		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		result := db.Create(&user)

		fmt.Println(result)

	}

	return nil
}
