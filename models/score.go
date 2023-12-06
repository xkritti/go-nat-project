package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type EngScore struct {
	Name         string         `json:"name"`
	HashCid      string         `json:"hash_cid" gorm:"unique"`
	LevelRange   string         `json:"level_range"`
	Province     string         `json:"province"`
	Region       string         `json:"region"`
	ExamLocation string         `json:"exam_location"`
	ScorePerPart datatypes.JSON `json:"score_per_part"`
	TotalScore   float64        `json:"total_score"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdateAt     time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type EngScorePerPart struct {
	Expression float64 `json:"expression"`
	Reading    float64 `json:"reading"`
	Structure  float64 `json:"structure"`
	Vocabulary float64 `json:"vocabulary"`
}
