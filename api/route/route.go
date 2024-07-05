package routers

import (
	"github.com/adon988/go_api_example/api/controllers"
	"github.com/adon988/go_api_example/api/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var memberController controllers.MemberController
var authController controllers.AuthController

func GetRouter(r *gin.Engine) {

	//swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authGroup := r.Group("/auth").Use(middleware.CORSMiddleware())
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/register", authController.Register)
	}

	memberGroup := r.Group("/member")
	memberGroup.Use(middleware.JWTAuthMiddleware())
	memberGroup.Use(middleware.CORSMiddleware())
	{
		memberGroup.GET("/", memberController.GetMemberInfo)
		memberGroup.GET("/roles", memberController.FindMembersWithRoles)
		memberGroup.PATCH("/", memberController.UpdateMember)
		memberGroup.DELETE("/", memberController.DeleteMember)
	}
}
