package models

type Role struct {
	Id       int32  `gorm:"primaryKey;autoIncrement"`
	Title    string `gorm:"size:255"`
	RoleType string `gorm:"size:255"`
	Image    string `gorm:"longtext"`
}
