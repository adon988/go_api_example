package models

import (
	"time"

	"gorm.io/gorm"
)

type Unit struct {
	Id             string `gorm:"primaryKey;size:24"`
	Title          string `gorm:"size:255"`
	OrganizationId string `gorm:"size:24;index"`
	CourseId       string `gorm:"size:24;index"`
	Order          int32  `gorm:"size:10"`
	CreaterId      string `gorm:"size:24;index;comment:'only note the creater id, not the permission'"`
	Publish        int32  `gorm:"size:1;index;default:0;comment:'(0 private, 1 public)'"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
