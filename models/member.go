package models

type Member struct {
	ID   string `gorm:"primaryKey;size:16"`
	Name string `gorm:"size:255"`
	Age  int64  `gorm:"size:10"`
}
