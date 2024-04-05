package main

import (
	"fmt"
	"math/rand"
	"strconv"

	migrations "github.com/adon988/go_api_example/migrations"
	model "github.com/adon988/go_api_example/models"
	"github.com/adon988/go_api_example/utils"
)

func init() {
	err := utils.InitConfig()
	if err != nil {
		panic("init config error:" + err.Error())
	}
	//Auto migrations
	migrations.AutoMigrations()
}

var InfoDb utils.InfoDb

func main() {

	//example for connect to db
	Db, err := InfoDb.InitDB()
	if err != nil {
		panic("init db err" + err.Error())
	}
	fmt.Println("addr", &Db, Db)

	// Generate a random number between 10000 and 99999
	randNum := rand.Intn(90000) + 10000
	// Convert the random number to a string
	randStr := strconv.Itoa(randNum)
	email := "jinzhu@gmail.com"
	member := model.Member{ID: randStr, Name: "Jinzhu", Age: 18, Email: &email}

	// Create a new record
	result := Db.Create(&member)

	if result.Error != nil {
		panic("create member error:" + result.Error.Error())
	}

	//Query the record
	var members = model.Member{ID: member.ID}
	Db.Find(&members)
	fmt.Println(members)

	//Uupdate the record
	members.Name = "Jinzhu2"
	result = Db.Save(&members)
	fmt.Println(result.RowsAffected)
	fmt.Println(members)

	//Soft delete the record
	Db.Delete(&members)

}
