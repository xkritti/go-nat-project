package utils

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"go-nat-project/models"
	"io"
	"os"

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

func ExcelReader(filename string, index int) (*excelize.File, string, int, error) {
	excelResult, err := excelize.OpenFile(filename)
	sheetName := excelResult.GetSheetName(index)
	if err != nil {
		return nil, "", 0, err
	}

	file, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, "", 0, err
	}

	rows, err := file.GetRows(file.GetSheetName(index))
	if err != nil {
		return nil, "", 0, err
	}

	return excelResult, sheetName, len(rows), nil
}

func UploadFileReader(c *fiber.Ctx) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", errors.New("Uploading Failed!")
	}

	src, err := file.Open()
	if err != nil {
		return "", errors.New("Source Invalid")
	}

	defer src.Close()

	dest, err := os.Create(file.Filename)

	defer dest.Close()

	_, err = io.Copy(dest, src)

	if err != nil {
		return "", err
	}

	fmt.Println("Upload Excel")

	return file.Filename, nil
}

func DeleteFile(filename string) error {
	fmt.Println(filename)
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}
