package models

import (
	"time"

	"gorm.io/gorm"
)

type Organization struct {
	Id              string `gorm:"primaryKey;size:24"`
	MemberId        string `gorm:"size:24;index"`
	Title           string `gorm:"size:255"`
	Order           int32  `gorm:"size:10"`
	SourceLangeuage string `gorm:"size:255"`
	TargetLanguage  string `gorm:"size:255"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
