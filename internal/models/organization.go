package models

import (
	"time"

	"gorm.io/gorm"
)

type Organization struct {
	Id             string         `gorm:"primaryKey;size:24"`
	Title          string         `gorm:"size:255"`
	Order          int32          `gorm:"size:10"`
	SourceLanguage string         `gorm:"size:255"`
	TargetLanguage string         `gorm:"size:255"`
	Publish        int32          `gorm:"size:1;index;default:0;comment:'(0 private, 1 public)'"`
	CreaterId      string         `gorm:"size:24;index;comment:'only note the creater id, not the permission'"`
	CreatedAt      time.Time      // Automatically managed by GORM for creation time
	UpdatedAt      time.Time      // Automatically managed by GORM for update time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
