package models

import (
	"time"

	"gorm.io/gorm"
)

var PermissionRoleEnum = map[string]int32{
	"admin":  1,
	"editor": 2,
	"viewer": 3,
}

// Permission represents the permissions a member has on various entities.
type OrganizationPermission struct {
	Id        string         `gorm:"primaryKey;size:24"`
	MemberId  string         `gorm:"size:24;index"`
	EntityId  string         `gorm:"size:24;index"`
	Role      int32          `gorm:"size:10"` // Role type: {admin:1, editor:2, guest:3}
	CreatedAt time.Time      // Automatically managed by GORM for creation time
	UpdatedAt time.Time      // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Automatically managed by GORM for soft delete
}

// Permission represents the permissions a member has on various entities.
type CoursePermission struct {
	Id        string         `gorm:"primaryKey;size:24"`
	MemberId  string         `gorm:"size:24;index"`
	EntityId  string         `gorm:"size:24;index"`
	Role      int32          `gorm:"size:10"` // Role type: {admin:1, editor:2, guest:3}
	CreatedAt time.Time      // Automatically managed by GORM for creation time
	UpdatedAt time.Time      // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Automatically managed by GORM for soft delete
}

// Permission represents the permissions a member has on various entities.
type UnitPermission struct {
	Id        string         `gorm:"primaryKey;size:24"`
	MemberId  string         `gorm:"size:24;index"`
	EntityId  string         `gorm:"size:24;index"`
	Role      int32          `gorm:"size:10"` // Role type: {admin:1, editor:2, guest:3}
	CreatedAt time.Time      // Automatically managed by GORM for creation time
	UpdatedAt time.Time      // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Automatically managed by GORM for soft delete
}
