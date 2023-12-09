package utils

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"go-nat-project/models"
	"io"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
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
	ID                string = "id"
	ExamType          string = "exam_type"
	Name              string = "name"
	Cid               string = "cid"
	NewCid            string = "new_cid"
	LevelRange        string = "level_range"
	Level             string = "level"
	Province          string = "province"
	Region            string = "region"
	School            string = "school"
	ExamLocation      string = "exam_location"
	TotalScore        string = "total_score"
	ProvinceRank      string = "province_rank"
	RegionRank        string = "region_rank"
	EngPtExpression   string = "score_pt_expression"
	EngPtReading      string = "score_pt_reading"
	EngPtStructure    string = "score_pt_structure"
	EngPtVocabulary   string = "score_pt_vocabulary"
	SciPtLesson       string = "score_pt_lesson_sci"
	SciPtApplied      string = "score_pt_applied_sci"
	MathPtCalculate   string = "score_pt_calculate_math"
	MathPtProblemMath string = "score_pt_problem_math"
	MathPtApplied     string = "score_pt_applied_math"

	MaxScore string = "max_score"
	MinScore string = "min_score"
	AvgScore string = "avg_score"

	NumberOfCompetitor string = "number_of_competitor"
	Year               string = "year"
	Subject            string = "subject"
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
		case SciPtLesson:
			columns[SciPtLesson] = string(alphabet[i])
			dbColumn = append(dbColumn, SciPtLesson)
		case SciPtApplied:
			columns[SciPtApplied] = string(alphabet[i])
			dbColumn = append(dbColumn, SciPtApplied)
		case MathPtCalculate:
			columns[MathPtCalculate] = string(alphabet[i])
			dbColumn = append(dbColumn, MathPtCalculate)
		case MathPtProblemMath:
			columns[MathPtProblemMath] = string(alphabet[i])
			dbColumn = append(dbColumn, MathPtProblemMath)
		case MathPtApplied:
			columns[MathPtApplied] = string(alphabet[i])
			dbColumn = append(dbColumn, MathPtApplied)
		case ProvinceRank:
			columns[ProvinceRank] = string(alphabet[i])
			dbColumn = append(dbColumn, ProvinceRank)
		case RegionRank:
			columns[RegionRank] = string(alphabet[i])
			dbColumn = append(dbColumn, RegionRank)
		case MaxScore:
			columns[MaxScore] = string(alphabet[i])
			dbColumn = append(dbColumn, MaxScore)
		case MinScore:
			columns[MinScore] = string(alphabet[i])
			dbColumn = append(dbColumn, MinScore)
		case AvgScore:
			columns[AvgScore] = string(alphabet[i])
			dbColumn = append(dbColumn, AvgScore)
		case NumberOfCompetitor:
			columns[NumberOfCompetitor] = string(alphabet[i])
			dbColumn = append(dbColumn, NumberOfCompetitor)
		case Year:
			columns[Year] = string(alphabet[i])
			dbColumn = append(dbColumn, Year)
		case Subject:
			columns[Subject] = string(alphabet[i])
			dbColumn = append(dbColumn, Subject)
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

func GetShortLevelRange(levelRangeTH string) (string, error) {
	switch levelRangeTH {
	case "ประถมศึกษาตอนต้น":
		return "LE", nil
	case "ประถมศึกษาตอนปลาย":
		return "UE", nil
	case "มัธยมศึกษาตอนต้น":
		return "LS", nil
	case "มัธยมศึกษาตอนปลาย":
		return "US", nil
	default:
		return "", errors.New("level range not found")
	}
}

func RemoveScientificNotationInString(s string) (string, error) {
	floatingPointNum, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return "", err
	}
	nonScientificNotation := strconv.FormatFloat(floatingPointNum, 'f', -1, 64)

	return nonScientificNotation, nil
}

func Validator() *validator.Validate {
	validate := validator.New()
	return validate
}
