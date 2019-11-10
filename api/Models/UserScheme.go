package Models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UUID      string `json:"uuid" gorm:"primary_key" validate:"omitempty|uuid"`
	FirstName string `json:"first_name" validate:"required|alpha|min=2"`
	LastName  string `json:"last_name" validate:"required|alph|min=2"`
	Email     string `json:"email" validate:"required|email"`
	Password  string `json:"password" validate:"required"`
	BirthDate time.Time `json:"birth_date" validate:"required"`
  AccessLevel int       `json:"access_level"`
}

func (b *User) TableName() string {
	return "user"
}
