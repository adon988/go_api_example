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
var organizationController controllers.OrganizationController
var courseController controllers.CourseController
var unitController controllers.UnitController

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
		memberGroup.PATCH("/", memberController.UpdateMember)
		memberGroup.DELETE("/", memberController.DeleteMember)

	}
	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.JWTAuthMiddleware())
	adminGroup.Use(middleware.CORSMiddleware())
	{
		adminGroup.GET("/organization", organizationController.GetOrganization)
		adminGroup.POST("/organization", organizationController.CreateOrganization)
		adminGroup.PATCH("/organization", organizationController.UpdateOrganization)
		adminGroup.DELETE("/organization", organizationController.DeleteOrganization)
		adminGroup.POST("/organization/assign", organizationController.AssignOrganizationPermission)

		adminGroup.GET("/course", courseController.GetCourse)
		adminGroup.POST("/course", courseController.CreateCourse)
		adminGroup.PATCH("/course", courseController.UpdateCourse)
		adminGroup.DELETE("/course", courseController.DeleteCourse)
		adminGroup.POST("/course/assign", courseController.AssignCoursePermission)

		adminGroup.GET("/units", unitController.GetUnits)
		adminGroup.POST("/unit", unitController.CreateUnit)
		adminGroup.PATCH("/unit", unitController.UpdateUnit)
		adminGroup.DELETE("/unit", unitController.DeleteUnit)
		adminGroup.POST("/unit/assign", unitController.AssignUnitPermission)
	}

}
