package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	Id        string         `gorm:"primaryKey;size:24"`
	Name      *string        `gorm:"size:255"` // A pointer to a string, allowing for null values
	Email     *string        `gorm:"size:255"`
	Birthday  *time.Time     // A pointer to a time.Time, allowing for null values
	Gender    *int           `gorm:"size:1"`
	CreatedAt time.Time      // Automatically managed by GORM for creation time
	UpdatedAt time.Time      // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Automatically managed by GORM for soft delete
}
