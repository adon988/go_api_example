package migrations

import (
	model "github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/utils"
)

var InfoDb utils.InfoDb

func AutoMigrations() {
	Db, err := InfoDb.InitDB()
	if err != nil {
		panic("init db err" + err.Error())
	}
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Authentication{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Organization{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Unit{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Word{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Count{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Course{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Role{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Member{})
}
