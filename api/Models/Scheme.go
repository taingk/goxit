package Models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UUID      string `json:"uuid" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
}

func (b *User) TableName() string {
	return "user"
}
