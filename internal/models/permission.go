package models

import (
	"time"

	"gorm.io/gorm"
)

// Permission represents the permissions a member has on various entities.
type OriganizationPermission struct {
	Id        string         `gorm:"primaryKey;size:24"`
	MemberId  string         `gorm:"size:24;index"`
	EntityId  string         `gorm:"size:24;index"`
	Role      string         `gorm:"size:50"` // Role type: admin, editor, view
	CreatedAt time.Time      // Automatically managed by GORM for creation time
	UpdatedAt time.Time      // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Automatically managed by GORM for soft delete
}

// Permission represents the permissions a member has on various entities.
type CoursePermission struct {
	Id        string         `gorm:"primaryKey;size:24"`
	MemberId  string         `gorm:"size:24;index"`
	EntityId  string         `gorm:"size:24;index"`
	Role      string         `gorm:"size:50"` // Role type: admin, editor, view
	CreatedAt time.Time      // Automatically managed by GORM for creation time
	UpdatedAt time.Time      // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Automatically managed by GORM for soft delete
}

// Permission represents the permissions a member has on various entities.
type UnitPermission struct {
	Id        string         `gorm:"primaryKey;size:24"`
	MemberId  string         `gorm:"size:24;index"`
	EntityId  string         `gorm:"size:24;index"`
	Role      string         `gorm:"size:50"` // Role type: admin, editor, view
	CreatedAt time.Time      // Automatically managed by GORM for creation time
	UpdatedAt time.Time      // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Automatically managed by GORM for soft delete
}
