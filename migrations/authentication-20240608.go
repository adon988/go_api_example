package migrations

import (
	"gorm.io/gorm"

	model "github.com/adon988/go_api_example/models"
)

func AuthenticationTableUp(Db *gorm.DB) {
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Authentication{})
}
