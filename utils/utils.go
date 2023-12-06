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

type ColumnLabelType string

const (
	ID           string = "id"
	ExamType     string = "exam_type"
	Name         string = "name"
	Cid          string = "cid"
	NewCid       string = "new_cid"
	LevelRange   string = "level_range"
	Level        string = "level"
	Province     string = "province"
	Region       string = "region"
	School       string = "school"
	ExamLocation string = "exam_location"
	TotalScore   string = "total_score"

	EngPtExpression string = "expression"
	EngPtReading    string = "reading"
	EngPtStructure  string = "structure"
	EngPtVocabulary string = "vocabulary"
)

type ColumnReader struct {
	Columns           map[string]string
	DBColumnsSelected []string
}

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewColumnReader(headerRow []string) *ColumnReader {
	fmt.Println(headerRow)
	columns := make(map[string]string)
	dbColumn := []string{}
	for i, v := range headerRow {
		switch v {
		case ID:
			columns[ID] = string(alphabet[i])
			dbColumn = append(dbColumn, ID)
		case ExamType:
			columns[ExamType] = string(alphabet[i])
			dbColumn = append(dbColumn, ExamType)
		case Name:
			columns[Name] = string(alphabet[i])
			dbColumn = append(dbColumn, Name)
		case Cid:
			columns[Cid] = string(alphabet[i])
			dbColumn = append(dbColumn, Cid)
		case NewCid:
			columns[NewCid] = string(alphabet[i])
			dbColumn = append(dbColumn, Cid)
		case LevelRange:
			columns[LevelRange] = string(alphabet[i])
			dbColumn = append(dbColumn, LevelRange)
		case Level:
			columns[Level] = string(alphabet[i])
			dbColumn = append(dbColumn, Level)
		case Province:
			columns[Province] = string(alphabet[i])
			dbColumn = append(dbColumn, Province)
		case Region:
			columns[Region] = string(alphabet[i])
			dbColumn = append(dbColumn, Region)
		case School:
			columns[School] = string(alphabet[i])
			dbColumn = append(dbColumn, School)
		case ExamLocation:
			columns[ExamLocation] = string(alphabet[i])
			dbColumn = append(dbColumn, ExamLocation)
		case TotalScore:
			columns[TotalScore] = string(alphabet[i])
			dbColumn = append(dbColumn, TotalScore)
		case EngPtExpression:
			columns[EngPtExpression] = string(alphabet[i])
			dbColumn = append(dbColumn, EngPtExpression)
		case EngPtReading:
			columns[EngPtReading] = string(alphabet[i])
			dbColumn = append(dbColumn, EngPtReading)
		case EngPtStructure:
			columns[EngPtStructure] = string(alphabet[i])
			dbColumn = append(dbColumn, EngPtStructure)
		case EngPtVocabulary:
			columns[EngPtVocabulary] = string(alphabet[i])
			dbColumn = append(dbColumn, EngPtVocabulary)
		}
	}
	fmt.Println(columns)
	return &ColumnReader{
		Columns:           columns,
		DBColumnsSelected: dbColumn,
	}
}

func (c *ColumnReader) GetColumnId(columnName string) string {
	return c.Columns[columnName]
}
