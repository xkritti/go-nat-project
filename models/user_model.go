package user_models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key, type:uuid; default:uuid_generate_v4()"`
	Username  string    `json:"username" gore:"unique"`
	Password  string    `json:"password"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email" gorm:"unique"`
	Phone     string    `json:"phone" goem:"unique"`
	// Role      string    `json:"role"`
}

type Users struct {
	Users []User `json:"users"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	return
}
