package Models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Vote struct {
	gorm.Model
	UUID        string         `json:"uuid" gorm:"primary_key"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	UUIDVotes   pq.StringArray `json:"uuid_votes" gorm:"type:text[]"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     time.Time      `json:"end_date"`
}

func (b *Vote) TableName() string {
	return "vote"
}
