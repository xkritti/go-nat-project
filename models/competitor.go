package models

import (
	"time"

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
