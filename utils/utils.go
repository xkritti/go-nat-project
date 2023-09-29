package utils

import (
	"crypto/sha256"
	"fmt"
	"go-nat-project/models"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

func SendSuccess(c *fiber.Ctx, data interface{}) error {
	return c.Status(200).JSON(models.CommonResponse{
		Code: 1000,
		Data: data,
	})
}

func SendCommonError(c *fiber.Ctx, errorData models.CommonError) error {
	return c.Status(200).JSON(errorData)
}

func GetSha256Enc(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	bs := h.Sum(nil)
	result := fmt.Sprintf("%x", bs)
	return result
}

func ExcelReader(filename string, index int) (*excelize.File, int, string, error) {
	excelResult, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, 0, "", err
	}

	file, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, 0, "", err
	}

	rows, err := file.GetRows(file.GetSheetName(index))
	if err != nil {
		return nil, 0, "", err
	}

	return excelResult, len(rows), "", nil
}
