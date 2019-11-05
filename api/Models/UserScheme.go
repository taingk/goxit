package Models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UUID        string    `json:"uuid" gorm:"primary_key"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	BirthDate   time.Time `json:"birth_date"`
	AccessLevel int       `json:"access_level"`
}

func (b *User) TableName() string {
	return "user"
}
