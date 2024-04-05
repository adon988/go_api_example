package migrations

import (
	model "go_api_example/models"
	"go_api_example/utils"
)

var InfoDb utils.InfoDb

func MembersTableUp() {

	Db, err := InfoDb.InitDB()
	if err != nil {
		panic("init db err" + err.Error())
	}
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Member{})
}