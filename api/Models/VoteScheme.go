package Models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Vote struct {
	gorm.Model
	UUID        string         `json:"uuid" gorm:"primary_key"`
	Title       string         `json:"title" gorm:"unique" validate:"required,min=2"`
	Description string         `json:"description" validate:"required,min=2"`
	UUIDVotes   pq.StringArray `json:"uuid_votes" gorm:"type:text[]"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     time.Time      `json:"end_date"`
	IsDeleted	int		  	   `json:"is_deleted"`
}

func (b *Vote) TableName() string {
	return "vote"
}
