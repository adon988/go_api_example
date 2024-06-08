package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	Id        string `gorm:"primaryKey;size:24"`
	Username  string `gorm:"size:255;index:idx_username"`
	Password  []byte
	Name      *string        `gorm:"size:255"` // A pointer to a string, allowing for null values
	Age       *int64         `gorm:"size:10"`
	Email     *string        `gorm:"size:255"`
	Address   *string        `gorm:"size:500"`
	CreatedAt time.Time      // Automatically managed by GORM for creation time
	UpdatedAt time.Time      // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Automatically managed by GORM for soft delete
}
