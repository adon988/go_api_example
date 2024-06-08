package migrations

import "github.com/adon988/go_api_example/utils"

var InfoDb utils.InfoDb

func AutoMigrations() {
	Db, err := InfoDb.InitDB()
	if err != nil {
		panic("init db err" + err.Error())
	}
	MembersTableUp(Db)
	AuthenticationTableUp(Db)
}
