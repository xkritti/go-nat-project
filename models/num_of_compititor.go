package models

import (
	"time"

	"gorm.io/gorm"
)

type NumberOfCompetitorByProvince struct {
	Year               string         `json:"year" gorm:"primaryKey"`
	LevelRange         string         `json:"level_range" gorm:"primaryKey"`
	Province           string         `json:"province" gorm:"primaryKey"`
	Subject            string         `json:"subject" gorm:"primaryKey"`
	NumberOfCompetitor int64          `json:"number_of_competitor"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at"`
}

type NumberOfCompetitorByRegion struct {
	Year               string         `json:"year" gorm:"primaryKey"`
	LevelRange         string         `json:"level_range" gorm:"primaryKey"`
	Region             string         `json:"region" gorm:"primaryKey"`
	Subject            string         `json:"subject" gorm:"primaryKey"`
	NumberOfCompetitor int64          `json:"number_of_competitor"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at"`
}
