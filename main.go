package main

import (
	"fmt"
	model "go_api_example/models"
	"go_api_example/utils"
)

func init() {
	err := utils.InitConfig()
	if err != nil {
		panic("init config error:" + err.Error())
	}
}

var InfoDb utils.InfoDb

func main() {

	//example for connect to db
	Db, err := InfoDb.InitDB()
	if err != nil {
		panic("init db err" + err.Error())
	}
	fmt.Println("addr", &Db, Db)
	var members = model.Member{ID: "1"}
	Db.Find(&members)
	fmt.Println(members)
}
