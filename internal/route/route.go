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
var wordController controllers.WordController
var quizController controllers.QuizController

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

	myGroup := r.Group("/my")
	myGroup.Use(middleware.JWTAuthMiddleware())
	myGroup.Use(middleware.CORSMiddleware())
	{
		myGroup.GET("/organization", organizationController.GetOrganizationMemberBelongTo)
		myGroup.POST("/organization", organizationController.CreateOrganization)
		myGroup.PATCH("/organization", organizationController.UpdateOrganization)
		myGroup.DELETE("/organization", organizationController.DeleteOrganization)
		myGroup.POST("/organization/assign", organizationController.AssignOrganizationPermission)

		myGroup.GET("/course", courseController.GetCourseMmeberBelongTo)
		myGroup.POST("/course", courseController.CreateCourse)
		myGroup.PATCH("/course", courseController.UpdateCourse)
		myGroup.DELETE("/course", courseController.DeleteCourse)
		myGroup.POST("/course/assign", courseController.AssignCoursePermission)

		myGroup.GET("course/:course_id/units", unitController.GetUnitsByCourseID)
		myGroup.GET("/units", unitController.GetUnitMemberBelongTo)
		myGroup.POST("/unit", unitController.CreateUnit)
		myGroup.PATCH("/unit", unitController.UpdateUnit)
		myGroup.DELETE("/unit", unitController.DeleteUnit)
		myGroup.POST("/unit/assign", unitController.AssignUnitPermission)

		myGroup.GET("/unit/:unit_id/words", wordController.GetWordsByUnitID)
		myGroup.POST("/word", wordController.CreateWord)
		myGroup.PATCH("/word", wordController.UpdateWord)
		myGroup.DELETE("/word", wordController.DeleteWord)

		myGroup.POST("/quiz", quizController.CreateQuiz)
		myGroup.GET("/quiz_list", quizController.GetQuizsListWithAnswersByMember)
		myGroup.GET("/quiz/:quiz_id", quizController.GetQuizByMember)
		myGroup.PATCH("/quiz_answer_record", quizController.UpdateQuizAnswerRecord)

	}

}
