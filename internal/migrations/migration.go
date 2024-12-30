package migrations

import (
	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/utils"
	casbin "github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var InfoDb utils.InfoDb
var Casbin *casbin.Enforcer

type NewMigration struct {
	Db *gorm.DB
}

func (r NewMigration) AutoMigration() {
	for _, model := range models.ModelsMap {
		r.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(model)

	}
}

func (r NewMigration) AutoMigrateCasbin() {
	model_config := utils.Configs.ACM.Model.Config
	a, err := gormadapter.NewAdapterByDB(r.Db)
	if err != nil {
		panic("error: adapter:" + err.Error())
	}
	casbin.NewEnforcer(model_config, a)
}

func (r NewMigration) MigrationTable(tableName string) bool {
	if model, ok := models.ModelsMap[tableName]; ok {
		r.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(model)
		return true
	}
	return false
}

func (r NewMigration) AlertTable(alertName string) bool {
	var err error
	if alertName == "rename_quiz_answer_records_column" {
		err = RenameCorrectAnswerCountToFailedAnswerCount(r.Db)
	}

	if err != nil {
		return false
	}

	return true
}
