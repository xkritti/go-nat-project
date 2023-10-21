package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              uuid.UUID `json:"id" gorm:"primary_key, type:uuid default:uuid_generate_v4()"`
	Cid             string    `json:"cid" gorm:"unique"`
	Prefix          string    `json:"prefix"`
	Name            string    `json:"name"`
	LevelType       string    `json:"level_type"`
	CompetitionType string    `json:"competition_type"`
	ExamLocation    string    `json:"exam_location"`
	School          string    `json:"school"`
	Province        string    `json:"province"`
	Sector          string    `json:"sector"`
	Level           string    `json:"level"`
}

type Users struct {
	Users []User `json:"users"`
}

// func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	// UUID version 4
// 	user.ID = uuid.New()
// 	return
// }
