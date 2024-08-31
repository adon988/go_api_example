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
	MemberId  string         `gorm:"primaryKey;size:24"`
	EntityId  string         `gorm:"primaryKey;size:24"`
	Role      int32          `gorm:"size:10"` // Role type: {admin:1, editor:2, guest:3}
	CreatedAt time.Time      // Automatically managed by GORM for creation time
	UpdatedAt time.Time      // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Automatically managed by GORM for soft delete
}

// Permission represents the permissions a member has on various entities.
type CoursePermission struct {
	MemberId  string         `gorm:"primaryKey;size:24"`
	EntityId  string         `gorm:"primaryKey;size:24"`
	Role      int32          `gorm:"size:10"` // Role type: {admin:1, editor:2, guest:3}
	CreatedAt time.Time      // Automatically managed by GORM for creation time
	UpdatedAt time.Time      // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Automatically managed by GORM for soft delete
}

// Permission represents the permissions a member has on various entities.
type UnitPermission struct {
	MemberId  string         `gorm:"primaryKey;size:24"`
	EntityId  string         `gorm:"primaryKey;size:24"`
	Role      int32          `gorm:"size:10"` // Role type: {admin:1, editor:2, guest:3}
	CreatedAt time.Time      // Automatically managed by GORM for creation time
	UpdatedAt time.Time      // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Automatically managed by GORM for soft delete
}
