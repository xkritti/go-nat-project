package models

import (
	"time"

	"gorm.io/gorm"
)

type Score struct {
	Name         string         `json:"name"`
	HashCid      string         `json:"hash_cid" gorm:"unique"`
	LevelRange   string         `json:"level_range"`
	Province     string         `json:"province"`
	Region       string         `json:"region"`
	ExamLocation string         `json:"exam_location"`
	TotalScore   float64        `json:"total_score"`
	ProvinceRank int            `json:"province_rank"`
	RegionRank   int            `json:"region_rank"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type EngScorePerPart struct {
	ScorePtExpression float64 `json:"score_pt_expression"`
	ScorePtReading    float64 `json:"score_pt_reading"`
	ScorePtStructure  float64 `json:"score_pt_structure"`
	ScorePtVocabulary float64 `json:"score_pt_vocabulary"`
}

type EngScorePerUser struct {
	TotalScore      float64 `json:"score_percentage"`
	ProvinceRank    int     `json:"province_rank"`
	RegionRank      int     `json:"region_rank"`
	EngScorePerPart `json:"parts"`
}

type EngScore struct {
	Score
	EngScorePerPart
}

type SciScorePerPart struct {
	ScorePtLessonSci  float64 `json:"score_pt_lesson_sci"`
	ScorePtAppliedSci float64 `json:"score_pt_applied_sci"`
}

type SciScorePerUser struct {
	TotalScore      float64 `json:"score_percentage"`
	ProvinceRank    int     `json:"province_rank"`
	RegionRank      int     `json:"region_rank"`
	SciScorePerPart `json:"parts"`
}

type SciScore struct {
	Score
	SciScorePerPart
}

type MathScorePerPart struct {
	ScorePtCalculate   float64 `json:"score_pt_calculate"`
	ScorePtProblemMath float64 `json:"score_pt_problem_math"`
	ScorePtAppliedMath float64 `json:"score_pt_applied_math"`
}

type MathScorePerUser struct {
	TotalScore       float64 `json:"score_percentage"`
	ProvinceRank     int     `json:"province_rank"`
	RegionRank       int     `json:"region_rank"`
	MathScorePerPart `json:"parts"`
}

type MathScore struct {
	Score
	MathScorePerPart
}

type UserScore struct {
	Subjects `json:"subjects"`
}

type Subjects struct {
	*EngScorePerUser  `json:"eng"`
	*SciScorePerUser  `json:"sci"`
	*MathScorePerUser `json:"math"`
}
