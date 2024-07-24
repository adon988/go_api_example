package utils

import (
	casbin "github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var Casbin *casbin.Enforcer

// InitCasbin is a function to initialize casbin
func InitCasbin() (*casbin.Enforcer, error) {
	var InfoDb InfoDb
	model_config := Configs.ACM.Model.Config
	db, _ := InfoDb.InitDB()
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic("error: adapter:" + err.Error())
	}
	e, err := casbin.NewEnforcer(model_config, a)
	return e, err
}
