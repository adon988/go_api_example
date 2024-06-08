package main

import (
	router "github.com/adon988/go_api_example/api/route"
	migrations "github.com/adon988/go_api_example/migrations"
	"github.com/adon988/go_api_example/utils"
	"github.com/gin-gonic/gin"
)

var InfoDb utils.InfoDb

func init() {
	err := utils.InitConfig()
	if err != nil {
		panic("init config error:" + err.Error())
	}
	//Auto migrations
	migrations.AutoMigrations()

	//Init swagger
	utils.InitSwagger()

}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	r := gin.Default()
	router.GetRouter(r)

	r.Run(`:8080`)
}
