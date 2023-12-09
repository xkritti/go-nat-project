package handler

import (
	"fmt"
	"go-nat-project/database"
	"go-nat-project/models"
	"go-nat-project/utils"
	"log/slog"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm/clause"
)

func GetScore(c *fiber.Ctx) error {
	return nil
}

func UploadScore(c *fiber.Ctx) error {

	filename, err := utils.UploadFileReader(c)
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Update Failed with %v", err.Error()),
		})
	}

	defer utils.DeleteFile(filename)

	shortSubjectName := strings.ToUpper(c.FormValue("short_subject_name"))
	if shortSubjectName == "" && (shortSubjectName != "ENG" && shortSubjectName != "SCI" && shortSubjectName != "MATH") {
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

	errRowNumberList := []string{}
	rows, err := excelResult.GetRows(sheetName)
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Update Failed with %v", err.Error()),
		})
	}
	colReader := utils.NewColumnReader(rows[1])

	// ENG
	if shortSubjectName == "ENG" {
		engScoreList := []*models.EngScore{}
		for i := 3; i <= len(rows); i++ {
			userInfo, errRow, err := getUserInfo(excelResult, sheetName, colReader, i)
			if err != nil {
				slog.Error(err.Error())
				errRowNumberList = append(errRowNumberList, fmt.Sprintf("%v, with error %v", errRow, err.Error()))
				continue
			}

			scorePtExpressionStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.EngPtExpression), i))
			scorePtExpression, err := strconv.ParseFloat(scorePtExpressionStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtExpression = 0
			}

			scorePtReadingStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.EngPtReading), i))
			scorePtReading, err := strconv.ParseFloat(scorePtReadingStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtReading = 0
			}

			scorePtStructureStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.EngPtStructure), i))
			scorePtStructure, err := strconv.ParseFloat(scorePtStructureStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtStructure = 0
			}

			scorePtVocabularyStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.EngPtVocabulary), i))
			scorePtVocabulary, err := strconv.ParseFloat(scorePtVocabularyStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtVocabulary = 0
			}

			engScoreList = append(engScoreList, &models.EngScore{
				Score: *userInfo,
				EngScorePerPart: models.EngScorePerPart{
					ScorePtExpression: scorePtExpression,
					ScorePtReading:    scorePtReading,
					ScorePtStructure:  scorePtStructure,
					ScorePtVocabulary: scorePtVocabulary,
				},
			})
		}
		tx := db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(engScoreList, 500)
		if tx.Error != nil {
			return c.JSON(models.CommonResponse{
				Code: 1001,
				Data: fmt.Sprintf("Upload Failed with %v", tx.Error.Error()),
			})
		}
		return c.JSON(models.CommonResponse{
			Code: 1000,
			Data: struct {
				Message string
				ErrList []string
			}{
				Message: fmt.Sprintf("Upload Complete total row = %v, insert into db %v record", len(engScoreList), tx.RowsAffected),
				ErrList: errRowNumberList,
			},
		})
	}

	// SCI
	if shortSubjectName == "SCI" {
		sciScoreList := []*models.SciScore{}
		for i := 3; i <= len(rows); i++ {
			userInfo, errRow, err := getUserInfo(excelResult, sheetName, colReader, i)
			if err != nil {
				slog.Error(err.Error())
				errRowNumberList = append(errRowNumberList, fmt.Sprintf("%v, with error %v", errRow, err.Error()))
				continue
			}
			scorePtLessonStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.SciPtLesson), i))
			scorePtLesson, err := strconv.ParseFloat(scorePtLessonStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtLesson = 0
			}

			scorePtAppliedStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.SciPtApplied), i))
			scorePtApplied, err := strconv.ParseFloat(scorePtAppliedStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtApplied = 0
			}

			sciScoreList = append(sciScoreList, &models.SciScore{
				Score: *userInfo,
				SciScorePerPart: models.SciScorePerPart{
					ScorePtLessonSci:  scorePtLesson,
					ScorePtAppliedSci: scorePtApplied,
				},
			})
		}
		tx := db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(sciScoreList, 500)
		if tx.Error != nil {
			return c.JSON(models.CommonResponse{
				Code: 1001,
				Data: fmt.Sprintf("Upload Failed with %v", tx.Error.Error()),
			})
		}
		return c.JSON(models.CommonResponse{
			Code: 1000,
			Data: struct {
				Message string
				ErrList []string
			}{
				Message: fmt.Sprintf("Upload Complete total row = %v, insert into db %v record", len(sciScoreList), tx.RowsAffected),
				ErrList: errRowNumberList,
			},
		})
	}

	// MATH
	if shortSubjectName == "MATH" {
		mathScoreList := []*models.MathScore{}
		for i := 3; i <= len(rows); i++ {
			userInfo, errRow, err := getUserInfo(excelResult, sheetName, colReader, i)
			if err != nil {
				slog.Error(err.Error())
				errRowNumberList = append(errRowNumberList, fmt.Sprintf("%v, with error %v", errRow, err.Error()))
				continue
			}
			scorePtCalculateStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.MathPtCalculate), i))
			scorePtCalculate, err := strconv.ParseFloat(scorePtCalculateStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtCalculate = 0
			}

			scorePtProblemMathStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.MathPtProblemMath), i))
			scorePtProblemMath, err := strconv.ParseFloat(scorePtProblemMathStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtProblemMath = 0
			}

			scorePtAppliedStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.MathPtApplied), i))
			scorePtApplied, err := strconv.ParseFloat(scorePtAppliedStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtApplied = 0
			}

			mathScoreList = append(mathScoreList, &models.MathScore{
				Score: *userInfo,
				MathScorePerPart: models.MathScorePerPart{
					ScorePtCalculate:   scorePtCalculate,
					ScorePtProblemMath: scorePtProblemMath,
					ScorePtAppliedMath: scorePtApplied,
				},
			})
		}
		tx := db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(mathScoreList, 500)
		if tx.Error != nil {
			return c.JSON(models.CommonResponse{
				Code: 1001,
				Data: fmt.Sprintf("Upload Failed with %v", tx.Error.Error()),
			})
		}
		return c.JSON(models.CommonResponse{
			Code: 1000,
			Data: struct {
				Message string
				ErrList []string
			}{
				Message: fmt.Sprintf("Upload Complete total row = %v, insert into db %v record", len(mathScoreList), tx.RowsAffected),
				ErrList: errRowNumberList,
			},
		})
	}

	return c.JSON(models.CommonResponse{
		Code: 9999,
		Data: struct {
			Message string
		}{
			Message: "Upload Failed with unknown error",
		},
	})
}

func UpdateScore(c *fiber.Ctx) error {
	filename, err := utils.UploadFileReader(c)
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Update Failed with %v", err.Error()),
		})
	}

	defer utils.DeleteFile(filename)

	shortSubjectName := strings.ToUpper(c.FormValue("short_subject_name"))
	if shortSubjectName == "" && (shortSubjectName != "ENG" && shortSubjectName != "SCI" && shortSubjectName != "MATH") {
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

	errRowNumberList := []int{}
	rows, err := excelResult.GetRows(sheetName)
	if err != nil {
		return c.JSON(models.CommonResponse{
			Code: 1001,
			Data: fmt.Sprintf("Update Failed with %v", err.Error()),
		})
	}
	colReader := utils.NewColumnReader(rows[1])

	// ENG
	if shortSubjectName == "ENG" {
		engScoreList := []*models.EngScore{}
		for i := 3; i <= len(rows); i++ {
			userInfo, errRow, err := getUserInfoForUpdate(excelResult, sheetName, colReader, i)
			if err != nil {
				slog.Error(err.Error())
				errRowNumberList = append(errRowNumberList, errRow)
				continue
			}

			scorePtExpressionStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.EngPtExpression), i))
			scorePtExpression, err := strconv.ParseFloat(scorePtExpressionStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtExpression = -99
			}

			scorePtReadingStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.EngPtReading), i))
			scorePtReading, err := strconv.ParseFloat(scorePtReadingStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtReading = -99
			}

			scorePtStructureStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.EngPtStructure), i))
			scorePtStructure, err := strconv.ParseFloat(scorePtStructureStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtStructure = -99
			}

			scorePtVocabularyStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.EngPtVocabulary), i))
			scorePtVocabulary, err := strconv.ParseFloat(scorePtVocabularyStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtVocabulary = -99
			}

			engScoreList = append(engScoreList, &models.EngScore{
				Score: *userInfo,
				EngScorePerPart: models.EngScorePerPart{
					ScorePtExpression: scorePtExpression,
					ScorePtReading:    scorePtReading,
					ScorePtStructure:  scorePtStructure,
					ScorePtVocabulary: scorePtVocabulary,
				},
			})
		}

		dbUpdateErrorList := []models.EngScore{}
		for _, v := range engScoreList {
			updateColumn := []string{}
			if v.Name != "" {
				updateColumn = append(updateColumn, utils.Name)
			}
			if v.LevelRange != "" {
				updateColumn = append(updateColumn, utils.LevelRange)
			}
			if v.Province != "" {
				updateColumn = append(updateColumn, utils.Province)
			}
			if v.Region != "" {
				updateColumn = append(updateColumn, utils.Region)
			}
			if v.ExamLocation != "" {
				updateColumn = append(updateColumn, utils.ExamLocation)
			}
			if v.ProvinceRank != -99 {
				updateColumn = append(updateColumn, utils.ProvinceRank)
			}
			if v.RegionRank != -99 {
				updateColumn = append(updateColumn, utils.RegionRank)
			}
			if v.TotalScore != -99 {
				updateColumn = append(updateColumn, utils.TotalScore)
			}
			if v.EngScorePerPart.ScorePtExpression != -99 {
				updateColumn = append(updateColumn, utils.EngPtExpression)
			}
			if v.EngScorePerPart.ScorePtReading != -99 {
				updateColumn = append(updateColumn, utils.EngPtReading)
			}
			if v.EngScorePerPart.ScorePtStructure != -99 {
				updateColumn = append(updateColumn, utils.EngPtStructure)
			}
			if v.EngScorePerPart.ScorePtVocabulary != -99 {
				updateColumn = append(updateColumn, utils.EngPtVocabulary)
			}

			tx := db.Model(&models.EngScore{}).Where("hash_cid = ?", v.HashCid).Select(updateColumn).Updates(v)
			if tx.Error != nil {
				dbUpdateErrorList = append(dbUpdateErrorList, *v)
			}
			if tx.RowsAffected == 0 || tx.RowsAffected > 1 {
				dbUpdateErrorList = append(dbUpdateErrorList, *v)
			}
		}
		return c.JSON(models.CommonResponse{
			Code: 1000,
			Data: struct {
				Message       string
				ErrList       []int
				DbUpdateError []models.EngScore
			}{
				Message:       fmt.Sprintf("Update Complete total row = %v", len(engScoreList)),
				ErrList:       errRowNumberList,
				DbUpdateError: dbUpdateErrorList,
			},
		})
	}

	// SCI
	if shortSubjectName == "SCI" {
		sciScoreList := []*models.SciScore{}
		for i := 3; i <= len(rows); i++ {
			userInfo, errRow, err := getUserInfoForUpdate(excelResult, sheetName, colReader, i)
			if err != nil {
				slog.Error(err.Error())
				errRowNumberList = append(errRowNumberList, errRow)
			}
			scorePtLessonStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.SciPtLesson), i))
			scorePtLesson, err := strconv.ParseFloat(scorePtLessonStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtLesson = -99
			}

			scorePtAppliedStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.SciPtApplied), i))
			scorePtApplied, err := strconv.ParseFloat(scorePtAppliedStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtApplied = -99
			}

			sciScoreList = append(sciScoreList, &models.SciScore{
				Score: *userInfo,
				SciScorePerPart: models.SciScorePerPart{
					ScorePtLessonSci:  scorePtLesson,
					ScorePtAppliedSci: scorePtApplied,
				},
			})
		}
		dbUpdateErrorList := []models.SciScore{}
		for _, v := range sciScoreList {
			updateColumn := []string{}
			if v.Name != "" {
				updateColumn = append(updateColumn, utils.Name)
			}
			if v.LevelRange != "" {
				updateColumn = append(updateColumn, utils.LevelRange)
			}
			if v.Province != "" {
				updateColumn = append(updateColumn, utils.Province)
			}
			if v.Region != "" {
				updateColumn = append(updateColumn, utils.Region)
			}
			if v.ExamLocation != "" {
				updateColumn = append(updateColumn, utils.ExamLocation)
			}
			if v.ProvinceRank != -99 {
				updateColumn = append(updateColumn, utils.ProvinceRank)
			}
			if v.RegionRank != -99 {
				updateColumn = append(updateColumn, utils.RegionRank)
			}
			if v.TotalScore != -99 {
				updateColumn = append(updateColumn, utils.TotalScore)
			}
			if v.SciScorePerPart.ScorePtLessonSci != -99 {
				updateColumn = append(updateColumn, utils.EngPtExpression)
			}
			if v.SciScorePerPart.ScorePtAppliedSci != -99 {
				updateColumn = append(updateColumn, utils.EngPtReading)
			}

			tx := db.Model(&models.EngScore{}).Where("hash_cid = ?", v.HashCid).Select(updateColumn).Updates(v)
			if tx.Error != nil {
				dbUpdateErrorList = append(dbUpdateErrorList, *v)
			}
			if tx.RowsAffected == 0 || tx.RowsAffected > 1 {
				dbUpdateErrorList = append(dbUpdateErrorList, *v)
			}
		}
		return c.JSON(models.CommonResponse{
			Code: 1000,
			Data: struct {
				Message       string
				ErrList       []int
				DbUpdateError []models.SciScore
			}{
				Message:       fmt.Sprintf("Update Complete total row = %v", len(sciScoreList)),
				ErrList:       errRowNumberList,
				DbUpdateError: dbUpdateErrorList,
			},
		})
	}

	// MATH
	if shortSubjectName == "MATH" {
		mathScoreList := []*models.MathScore{}
		for i := 3; i <= len(rows); i++ {
			userInfo, errRow, err := getUserInfo(excelResult, sheetName, colReader, i)
			if err != nil {
				slog.Error(err.Error())
				errRowNumberList = append(errRowNumberList, errRow)
				continue
			}
			scorePtCalculateStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.MathPtCalculate), i))
			scorePtCalculate, err := strconv.ParseFloat(scorePtCalculateStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtCalculate = -99
			}

			scorePtProblemMathStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.MathPtProblemMath), i))
			scorePtProblemMath, err := strconv.ParseFloat(scorePtProblemMathStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtProblemMath = -99
			}

			scorePtAppliedStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.MathPtApplied), i))
			scorePtApplied, err := strconv.ParseFloat(scorePtAppliedStr, 64)
			if err != nil {
				slog.Error(err.Error())
				scorePtApplied = -99
			}

			mathScoreList = append(mathScoreList, &models.MathScore{
				Score: *userInfo,
				MathScorePerPart: models.MathScorePerPart{
					ScorePtCalculate:   scorePtCalculate,
					ScorePtProblemMath: scorePtProblemMath,
					ScorePtAppliedMath: scorePtApplied,
				},
			})
		}
		dbUpdateErrorList := []models.MathScore{}
		for _, v := range mathScoreList {
			updateColumn := []string{}
			if v.Name != "" {
				updateColumn = append(updateColumn, utils.Name)
			}
			if v.LevelRange != "" {
				updateColumn = append(updateColumn, utils.LevelRange)
			}
			if v.Province != "" {
				updateColumn = append(updateColumn, utils.Province)
			}
			if v.Region != "" {
				updateColumn = append(updateColumn, utils.Region)
			}
			if v.ExamLocation != "" {
				updateColumn = append(updateColumn, utils.ExamLocation)
			}
			if v.ProvinceRank != -99 {
				updateColumn = append(updateColumn, utils.ProvinceRank)
			}
			if v.RegionRank != -99 {
				updateColumn = append(updateColumn, utils.RegionRank)
			}
			if v.TotalScore != -99 {
				updateColumn = append(updateColumn, utils.TotalScore)
			}
			if v.MathScorePerPart.ScorePtCalculate != -99 {
				updateColumn = append(updateColumn, utils.EngPtExpression)
			}
			if v.MathScorePerPart.ScorePtProblemMath != -99 {
				updateColumn = append(updateColumn, utils.EngPtReading)
			}
			if v.MathScorePerPart.ScorePtAppliedMath != -99 {
				updateColumn = append(updateColumn, utils.EngPtStructure)
			}

			tx := db.Model(&models.EngScore{}).Where("hash_cid = ?", v.HashCid).Select(updateColumn).Updates(v)
			if tx.Error != nil {
				dbUpdateErrorList = append(dbUpdateErrorList, *v)
			}
			if tx.RowsAffected == 0 || tx.RowsAffected > 1 {
				dbUpdateErrorList = append(dbUpdateErrorList, *v)
			}
		}
		return c.JSON(models.CommonResponse{
			Code: 1000,
			Data: struct {
				Message       string
				ErrList       []int
				DbUpdateError []models.MathScore
			}{
				Message:       fmt.Sprintf("Update Complete total row = %v", len(mathScoreList)),
				ErrList:       errRowNumberList,
				DbUpdateError: dbUpdateErrorList,
			},
		})
	}

	return c.JSON(models.CommonResponse{
		Code: 9999,
		Data: struct {
			Message string
		}{
			Message: "Upload Failed with unknown error",
		},
	})
}

func getUserInfo(excelResult *excelize.File, sheetName string, colReader *utils.ColumnReader, row int) (userScore *models.Score, rowNumberOfError int, err error) {

	rowNumberOfError = row

	name, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Name), row))
	cid, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Cid), row), excelize.Options{RawCellValue: true})
	levelRange, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.LevelRange), row))
	province, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Province), row))
	region, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Region), row))
	examLocation, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.ExamLocation), row))
	totalScoreStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.TotalScore), row))
	totalScore, err := strconv.ParseFloat(totalScoreStr, 64)
	if err != nil {
		slog.Error(err.Error())
	}

	strCid, err := utils.RemoveScientificNotationInString(strings.TrimSpace(cid))
	if err != nil {
		strCid = strings.TrimSpace(cid)
	}

	if cid == "" {
		return nil, rowNumberOfError, fmt.Errorf("cid is empty")
	}

	provinceRankStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.ProvinceRank), row))
	provinceRank, err := strconv.Atoi(provinceRankStr)
	if err != nil {
		slog.Error(err.Error())
		provinceRank = 0
	}

	regionRankStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.RegionRank), row))
	regionRank, err := strconv.Atoi(regionRankStr)
	if err != nil {
		slog.Error(err.Error())
		regionRank = 0
	}

	shortLevelRange, err := utils.GetShortLevelRange(levelRange)
	if err != nil {
		slog.Error(err.Error())
		return nil, rowNumberOfError, err
	}

	fmt.Printf("| %v | %v | %v | %v | %v | %v | %v \n", name, cid, strCid, levelRange, province, region, examLocation)

	hashCid := utils.GetSha256Enc(strings.TrimSpace(strCid))
	userScore = &models.Score{
		Name:         name,
		HashCid:      hashCid,
		LevelRange:   shortLevelRange,
		Province:     province,
		Region:       region,
		ExamLocation: examLocation,
		TotalScore:   totalScore,
		ProvinceRank: provinceRank,
		RegionRank:   regionRank,
	}

	return userScore, 0, nil
}

func getUserInfoForUpdate(excelResult *excelize.File, sheetName string, colReader *utils.ColumnReader, row int) (userScore *models.Score, rowNumberOfError int, err error) {

	rowNumberOfError = row

	name, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Name), row))
	cid, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Cid), row), excelize.Options{RawCellValue: true})
	levelRange, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.LevelRange), row))
	province, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Province), row))
	region, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.Region), row))
	examLocation, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.ExamLocation), row))

	strCid, err := utils.RemoveScientificNotationInString(strings.TrimSpace(cid))
	if err != nil {
		strCid = strings.TrimSpace(cid)
	}

	if cid == "" {
		return nil, rowNumberOfError, fmt.Errorf("cid is empty")
	}

	totalScoreStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.TotalScore), row))
	totalScore, err := strconv.ParseFloat(totalScoreStr, 64)
	if err != nil {
		totalScore = -99
	}

	provinceRankStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.ProvinceRank), row))
	provinceRank, err := strconv.Atoi(provinceRankStr)
	if err != nil {
		provinceRank = -99
	}

	regionRankStr, _ := excelResult.GetCellValue(sheetName, fmt.Sprintf("%v%v", colReader.GetColumnId(utils.RegionRank), row))
	regionRank, err := strconv.Atoi(regionRankStr)
	if err != nil {
		regionRank = -99
	}

	shortLevelRange, err := utils.GetShortLevelRange(levelRange)
	if err != nil {
		shortLevelRange = ""
	}

	hashCid := utils.GetSha256Enc(strings.TrimSpace(strCid))
	userScore = &models.Score{
		Name:         name,
		HashCid:      hashCid,
		LevelRange:   shortLevelRange,
		Province:     province,
		Region:       region,
		ExamLocation: examLocation,
		TotalScore:   totalScore,
		ProvinceRank: provinceRank,
		RegionRank:   regionRank,
	}

	return userScore, 0, nil
}

func GetUserScore(c *fiber.Ctx) error {
	db := database.DB.Db
	competitor := &models.CompetitorWithScores{}

	query := `
SELECT ct.id ,ct.cid , ct.name , ct.exam_type , ct.level_range,

es.province_rank as eng_province_rank,
es.region_rank as eng_region_rank,

es.total_score as eng_total_score, 
es.score_pt_expression,
es.score_pt_reading,
es.score_pt_structure,
es.score_pt_vocabulary,

ss.province_rank as sci_province_rank,
ss.region_rank as sci_region_rank,

ss.total_score as sci_total_score, 
ss.score_pt_lesson_sci,
ss.score_pt_applied_sci,


ms.province_rank as math_province_rank,
ms.region_rank as math_region_rank,

ms.total_score as math_total_score ,  
ms.score_pt_calculate as score_pt_calculate_math,
ms.score_pt_problem_math,
ms.score_pt_applied_math

FROM competitors ct
LEFT JOIN eng_scores es ON ct.cid = es.hash_cid 
LEFT JOIN sci_scores ss ON ct.cid = ss.hash_cid
LEFT JOIN math_scores ms ON ct.cid = ms.hash_cid

WHERE ct.cid = '5f034889cfbb51ea2b309a62273b6c9fe3ebe87c6be8a8f2e38de35cc67bc498'
`
	err := db.Raw(query).Scan(competitor).Error

	if err != nil {
		return c.SendStatus(500)
	}

	result := &models.UserScore{
		CID:        competitor.CID,
		Name:       competitor.Name,
		ExamType:   competitor.ExamType,
		LevelRange: competitor.LevelRange,
		MathInfo: models.MathInfo{
			MathTotalScore:   competitor.MathTotalScore,
			MathProvinceRank: competitor.MathProvinceRank,
			MathRegionRank:   competitor.MathRegionRank,
			Parts: models.MathPart{
				ScorePtCalculateMath: competitor.ScorePtCalculateMath,
				ScorePtProblemMath:   competitor.ScorePtProblemMath,
				ScorePtAppliedMath:   competitor.ScorePtAppliedMath,
			},
		},
		SciInfo: models.SciInfo{
			SciProvinceRank: competitor.SciProvinceRank,
			SciTotalScore:   competitor.SciTotalScore,
			SciRegionRank:   competitor.SciRegionRank,
			Parts: models.SciPart{
				ScorePtLessonSci:  competitor.ScorePtLessonSci,
				ScorePtAppliedSci: competitor.ScorePtAppliedSci,
			},
		},
		EngInfo: models.EngInfo{
			EngProvinceRank: competitor.EngProvinceRank,
			EngRegionRank:   competitor.EngRegionRank,
			EngTotalScore:   competitor.EngTotalScore,
			Parts: models.EngPart{
				ScorePtExpression: competitor.ScorePtExpression,
				ScorePtReading:    competitor.ScorePtReading,
				ScorePtStructure:  competitor.ScorePtStructure,
				ScorePtVocabulary: competitor.ScorePtVocabulary,
			},
		},
	}

	return utils.SendSuccess(c, result)
}
