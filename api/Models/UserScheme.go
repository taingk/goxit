package Models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UUID      string `json:"uuid" gorm:"primary_key" validate:"omitempty,uuid"`
	Firstname string `json:"firstname" validate:"required,min=2"`
	Lastname  string `json:"lastname" validate:"required,min=2"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	// BirthDate   time.Time `json:"birth_date" validate:"required"`
	AccessLevel int `json:"access_level"`
}

func (b *User) TableName() string {
	return "user"
}
