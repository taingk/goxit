package Models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UUID        string    `json:"uuid" gorm:"primary_key" validate:"omitempty,uuid"`
	Firstname   string    `json:"firstname" validate:"required,min=2"`
	Lastname    string    `json:"lastname" validate:"required,min=2"`
	Email       string    `json:"email" gorm:"unique" validate:"required,email"`
	Password    string    `json:"password" validate:"required,min=6"`
	BirthDate   time.Time `json:"birth_date"`
	AccessLevel int       `json:"access_level"`
	IsDeleted   int       `json:"is_deleted"`
}

func (b *User) TableName() string {
	return "user"
}
