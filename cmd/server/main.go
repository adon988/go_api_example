package main

import (
	"fmt"
	"strconv"

	router "github.com/adon988/go_api_example/internal/route"
	"github.com/adon988/go_api_example/internal/utils"
	"github.com/gin-gonic/gin"
)

var InfoDb utils.InfoDb

func init() {
	err := utils.InitConfig()
	if err != nil {
		panic("init config error:" + err.Error())
	}

	//Init swagger
	utils.InitSwagger()

	//Init Casbin
	utils.Casbin, err = utils.InitCasbin()

	if err != nil {
		panic("error: model:" + err.Error())
	}

}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if utils.Configs.Gin.Debug_Mode {
		gin.SetMode(gin.DebugMode) //gin debug mode
	} else {
		gin.SetMode(gin.ReleaseMode) //gin release mode
	}
	r := gin.Default()
	r.Use(gin.Recovery()) // prevent gin panic
	r.Use(gin.Logger())   // gin default logger
	router.GetRouter(r)
	fmt.Println()
	port := utils.Configs.Gin.Port
	r.Run(":" + strconv.Itoa(port))
}
