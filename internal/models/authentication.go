package models

import (
	"time"

	"gorm.io/gorm"
)

type Authentication struct {
	Username  string `gorm:"size:255;index:idx_username"`
	Password  []byte
	MemberId  string         `gorm:"size:24;index"`
	Type      *string        `gorm:"size:255; default:'ApiKeyAuth'; comment:'ApikeyAuth, AppleId, GoogleId..'"` // A pointer to a string, allowing for null values
	CreatedAt time.Time      // Automatically managed by GORM for creation time
	UpdatedAt time.Time      // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Automatically managed by GORM for soft delete
}
