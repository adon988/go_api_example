package models

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	Id             string       `gorm:"primaryKey;size:24"`
	Title          string       `gorm:"size:255"`
	OrganizationId string       `gorm:"size:24;index"`
	Order          int32        `gorm:"size:10"`
	PublishStatus  int32        `gorm:"size:10;default:3;comment:'(1 public, 2 member_public, 3 draft)'"`
	CreaterId      string       `gorm:"size:24;index;comment:'only note the creater id, not the permission'"`
	Permissions    []Permission `gorm:"type:json"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
