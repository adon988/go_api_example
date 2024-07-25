package models

import (
	"time"

	"gorm.io/gorm"
)

type Organization struct {
	Id              string `gorm:"primaryKey;size:24"`
	Title           string `gorm:"size:255"`
	Order           int32  `gorm:"size:10"`
	SourceLangeuage string `gorm:"size:255"`
	TargetLanguage  string `gorm:"size:255"`
	CreaterId       string `gorm:"size:24;index;comment:'only note the creater id, not the permission'"`
	Permissions     string `gorm:"type:json;column:permissions"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
