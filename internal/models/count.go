package models

import (
	"time"

	"gorm.io/gorm"
)

type Count struct {
	Id             string  `gorm:"primaryKey;size:24"`
	OrganizationId *string `gorm:"size:24;index"`
	CourseId       *string `gorm:"size:24;index"`
	UnitId         *string `gorm:"size:24;index"`
	UnitCount      int32   `gorm:"size:10;default:0"`
	WordCount      int32   `gorm:"size:10;default:0"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
