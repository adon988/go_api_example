package routers

import (
	"github.com/adon988/go_api_example/internal/controllers"
	"github.com/adon988/go_api_example/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var memberController controllers.MemberController
var authController controllers.AuthController
var roleController controllers.RoleController
var organizationController controllers.OrganizationController

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

	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.JWTAuthMiddleware())
	adminGroup.Use(middleware.CORSMiddleware())
	{
		adminGroup.GET("/roles", roleController.GetRole)
		adminGroup.POST("/roles", roleController.CreateRole)
		adminGroup.PATCH("/roles", roleController.UpdateRole)
		adminGroup.DELETE("/roles", roleController.DeleteRole)
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
	orgGroup := r.Group("/admin")
	orgGroup.Use(middleware.JWTAuthMiddleware())
	orgGroup.Use(middleware.CORSMiddleware())
	{
		orgGroup.GET("/organization", organizationController.GetOrganization)
		orgGroup.POST("/organization", organizationController.CreateOrganization)
		orgGroup.PATCH("/organization", organizationController.UpdateOrganization)
		orgGroup.DELETE("/organization", organizationController.DeleteOrganization)
	}

}
