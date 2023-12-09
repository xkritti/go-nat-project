package models

import (
	"time"

	"gorm.io/gorm"
)

type AvgScoreBySubject struct {
	Year       string         `json:"year" gorm:"primaryKey"`
	LevelRange string         `json:"level_range" gorm:"primaryKey"`
	Subject    string         `json:"subject" gorm:"primaryKey"`
	AvgScore   float64        `json:"avg_score"`
	MaxScore   float64        `json:"max_score"`
	MinScore   float64        `json:"min_score"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
