package handler

import (
	"encoding/json"
	"fmt"
	"go-nat-project/database"
	"go-nat-project/models"
	"go-nat-project/utils"
	"log/slog"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm/clause"
)

func GetScore(c *fiber.Ctx) error {
	return nil
}

type UploadScoreRequest struct {
	ShortSubjectName string `json:"short_subject_name"`
	sheetIndex       int    `json:"sheet_index"`
}

func UploadScore(c *fiber.Ctx) error {

	filename, err := utils.UploadFileReader(c)
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Update Failed with %v", err.Error()),
		})
	}

	shortSubjectName := strings.ToUpper(c.FormValue("short_subject_name"))
	if shortSubjectName == "" && shortSubjectName != "ENG" && shortSubjectName != "SCI" && shortSubjectName != "MATH" {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Update Failed with %v", err.Error()),
		})
	}

	sheetIndex, err := strconv.Atoi(c.FormValue("sheet_index"))
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Fail to convert sheet index with %v", err.Error()),
		})
	}

	excelResult, sheetName, _, err := utils.ExcelReader(filename, sheetIndex)
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Update Failed with %v", err.Error()),
		})
	}

	db := database.DB.Db

	if shortSubjectName == "ENG" {
		scoreList := []*models.EngScore{}
		errList := []string{}

		rows, err := excelResult.GetRows(sheetName)
		if err != nil {
			return c.JSON(models.CommonResponse{
				Code: 1001,
				Data: fmt.Sprintf("Update Failed with %v", err.Error()),
			})
		}

		colReader := utils.NewColumnReader(rows[1])
		for i := 3; i <= len(rows); i++ {
			name, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Name), i))
			cid, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Cid), i), excelize.Options{RawCellValue: true})
			levelRange, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.LevelRange), i))
			province, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Province), i))
			region, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Region), i))
			examLocation, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.ExamLocation), i))
			totalScoreStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.TotalScore), i))

			scorePtExpressionStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.EngPtExpression), i))
			scorePtReadingStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.EngPtReading), i))
			scorePtStructureStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.EngPtStructure), i))
			scorePtVocabularyStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.EngPtVocabulary), i))

			fmt.Printf("| %v | %v | %v | %v | %v | %v | %v | %v | %v  | %v  | %v  \n", name, cid, levelRange, province, region, examLocation, totalScoreStr, scorePtExpressionStr, scorePtReadingStr, scorePtStructureStr, scorePtVocabularyStr)

			scorePtExpression, err := strconv.ParseFloat(scorePtExpressionStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtExpression = 0
			}
			scorePtReading, err := strconv.ParseFloat(scorePtReadingStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtReading = 0
			}
			scorePtStructure, err := strconv.ParseFloat(scorePtStructureStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtStructure = 0
			}
			scorePtVocabulary, err := strconv.ParseFloat(scorePtVocabularyStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtVocabulary = 0
			}
			totalScore, err := strconv.ParseFloat(totalScoreStr, 64)
			if err != nil {
				slog.Error(err.Error())
				totalScore = 0
			}

			hashCid := utils.GetSha256Enc(cid)

			shortLevelRange := ""

			if levelRange == "ประถมศึกษาตอนต้น" {
				shortLevelRange = "LE"
			} else if levelRange == "ประถมศึกษาตอนปลาย" {
				shortLevelRange = "UE"
			} else if levelRange == "มัธยมศึกษาตอนต้น" {
				shortLevelRange = "LS"
			} else if levelRange == "มัธยมศึกษาตอนปลาย" {
				shortLevelRange = "US"
			} else {
				slog.Error("level range not found")
				errList = append(errList, cid)
				continue
			}

			userScore := &models.EngScore{
				Name:         name,
				HashCid:      hashCid,
				LevelRange:   shortLevelRange,
				Province:     province,
				Region:       region,
				ExamLocation: examLocation,
				TotalScore:   totalScore,
			}

			scorePerPath := &models.EngScorePerPart{
				Expression: scorePtExpression,
				Reading:    scorePtReading,
				Structure:  scorePtStructure,
				Vocabulary: scorePtVocabulary,
			}

			scorePerPathByte, err := json.Marshal(scorePerPath)
			if err != nil {
				slog.Error(err.Error())
				errList = append(errList, cid)
				continue
			}
			userScore.ScorePerPart = datatypes.JSON(scorePerPathByte)

			scoreList = append(scoreList, userScore)

		}

		tx := db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(scoreList, 500)
		if tx.Error != nil {
			return c.JSON(models.CommonResponse{
				Code: 1001,
				Data: fmt.Sprintf("Upload Failed with %v", tx.Error.Error()),
			})
		}
		utils.DeleteFile(filename)

		return c.JSON(models.CommonResponse{
			Code: 1000,
			Data: struct {
				Message string
				ErrList []string
			}{
				Message: fmt.Sprintf("Upload Complete total row = %v, insert into db %v record", len(scoreList), tx.RowsAffected),
				ErrList: errList,
			},
		})

	}

	return nil
}
