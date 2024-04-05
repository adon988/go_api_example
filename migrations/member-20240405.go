package migrations

import (
	"github.com/adon988/go_api_example/utils"

	model "github.com/adon988/go_api_example/models"
)

var InfoDb utils.InfoDb

func MembersTableUp() {

	Db, err := InfoDb.InitDB()
	if err != nil {
		panic("init db err" + err.Error())
	}
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Member{})
}
