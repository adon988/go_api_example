package models

import (
	"time"

	"gorm.io/gorm"
)

type Word struct {
	Id            string `gorm:"primaryKey;size:24"`
	UnitId        string `gorm:"size:24"`
	Word          string `gorm:"size:255"`
	Definition    string `gorm:"longtext"`
	Image         string `gorm:"longtext"`
	Pronunciation string `gorm:"longtext"`
	Description   string `gorm:"longtext"`
	Comment       string `gorm:"longtext"`
	Order         int32  `gorm:"size:10"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
