package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Competitor struct {
	gorm.Model
	ID           uuid.UUID `json:"id" gorm:"primary_key, type:uuid default:uuid_generate_v4()"`
	ExamType     string    `json:"type"`
	Name         string    `json:"name"`
	Cid          string    `json:"cid" gorm:"unique"`
	LevelRange   string    `json:"level_range"`
	Level        string    `json:"level"`
	Province     string    `json:"province"`
	Region       string    `json:"region"`
	School       string    `json:"school"`
	ExamLocation string    `json:"exam_location"`
}

type Competitors struct {
	Competitors []Competitor `json:"competitors"`
}

func (competitor *Competitor) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	competitor.ID = uuid.New()
	return
}
