package routers

import (
	"github.com/adon988/go_api_example/api/controllers"
	"github.com/adon988/go_api_example/api/middleware"
	"github.com/gin-gonic/gin"
)

func GetRouter(r *gin.Engine) {

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/jwt_token", func(ctx *gin.Context) {
		token, _ := middleware.GenToken("account_id")
		ctx.JSON(200, gin.H{
			"message": "success",
			"token":   token,
		})
	})
	memberGroup := r.Group("/member")
	memberGroup.Use(middleware.JWTAuthMiddleware())
	memberGroup.Use(middleware.CORSMiddleware())
	{
		memberGroup.GET("/:id", new(controllers.MemberController).GetMmeber)
		memberGroup.POST("/:id", new(controllers.MemberController).CreateMember)
		memberGroup.DELETE("/:id", new(controllers.MemberController).DeleteMember)
	}
}
