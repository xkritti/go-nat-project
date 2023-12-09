package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Competitor struct {
	// gorm.Model
	ID           string         `json:"id" gorm:"primary_key"`
	ExamType     string         `json:"type"`
	Name         string         `json:"name"`
	Cid          string         `json:"cid" gorm:"unique"`
	LevelRange   string         `json:"level_range"`
	Level        string         `json:"level"`
	Province     string         `json:"province"`
	Region       string         `json:"region"`
	School       string         `json:"school"`
	ExamLocation string         `json:"exam_location"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdateAt     time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type Competitors struct {
	Competitors []Competitor `json:"competitors"`
}

// func (competitor *Competitor) BeforeCreate(tx *gorm.DB) (err error) {
// 	// UUID version 4
// 	competitor.ID = uuid.New()
// 	return
// }

type CompetitorWithScores struct {
	ID         uuid.UUID `gorm:"column:id"`
	CID        string    `gorm:"column:cid"`
	Name       string    `gorm:"column:name"`
	ExamType   string    `gorm:"column:exam_type"`
	LevelRange string    `gorm:"column:level_range"`

	MathTotalScore   float64 `gorm:"column:math_total_score"`
	MathProvinceRank float64 `gorm:"column:math_province_rank"`
	MathRegionRank   float64 `gorm:"column:math_region_rank"`

	ScorePtCalculateMath float64 `gorm:"column:score_pt_calculate_math"`
	ScorePtProblemMath   float64 `gorm:"column:score_pt_problem_math"`
	ScorePtAppliedMath   float64 `gorm:"column:score_pt_applied_math"`

	SciProvinceRank float64 `gorm:"column:sci_province_rank"`
	SciRegionRank   float64 `gorm:"column:sci_region_rank"`

	SciTotalScore     float64 `gorm:"column:sci_total_score"`
	ScorePtLessonSci  float64 `gorm:"column:score_pt_lesson_sci"`
	ScorePtAppliedSci float64 `gorm:"column:score_pt_applied_sci"`

	EngProvinceRank float64 `gorm:"column:eng_province_rank"`
	EngRegionRank   float64 `gorm:"column:eng_region_rank"`
	EngTotalScore   float64 `gorm:"column:eng_total_score"`

	ScorePtExpression float64 `gorm:"column:score_pt_expression"`
	ScorePtReading    float64 `gorm:"column:score_pt_reading"`
	ScorePtStructure  float64 `gorm:"column:score_pt_structure"`
	ScorePtVocabulary float64 `gorm:"column:score_pt_vocabulary"`
}

type UserScore struct {
	CID        string   `json:"cid"`
	Name       string   `json:"name"`
	ExamType   string   `json:"exam_type"`
	LevelRange string   `json:"level_range"`
	MathInfo   MathInfo `json:"math"`
	SciInfo    SciInfo  `json:"sci"`
	EngInfo    EngInfo  `json:"eng"`
}

type MathInfo struct {
	MathTotalScore   float64  `json:"total_score"`
	MathProvinceRank float64  `json:"province_rank"`
	MathRegionRank   float64  `json:"region_rank"`
	Parts            MathPart `json:"parts"`
}

type MathPart struct {
	ScorePtCalculateMath float64 `json:"calculate_pt"`
	ScorePtProblemMath   float64 `json:"problem_pt"`
	ScorePtAppliedMath   float64 `json:"applied_pt"`
}

type SciInfo struct {
	SciProvinceRank float64 `json:"province_rank"`
	SciRegionRank   float64 `json:"region_rank"`
	SciTotalScore   float64 `json:"total_score"`
	Parts           SciPart `json:"parts"`
}

type SciPart struct {
	ScorePtLessonSci  float64 `json:"lesson_pt"`
	ScorePtAppliedSci float64 `json:"applied_pt"`
}

type EngInfo struct {
	EngProvinceRank float64 `json:"province_rank"`
	EngRegionRank   float64 `json:"region_rank"`
	EngTotalScore   float64 `json:"total_score"`
	Parts           EngPart `json:"parts"`
}

type EngPart struct {
	ScorePtExpression float64 `json:"expression_pt"`
	ScorePtReading    float64 `json:"reading_pt"`
	ScorePtStructure  float64 `json:"struct_pt"`
	ScorePtVocabulary float64 `json:"vocabuary_pt"`
}
