package migrations

import (
	model "github.com/adon988/go_api_example/models"
	"github.com/adon988/go_api_example/utils"
)

var InfoDb utils.InfoDb

func AutoMigrations() {
	Db, err := InfoDb.InitDB()
	if err != nil {
		panic("init db err" + err.Error())
	}
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Authentication{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Member{})

}
