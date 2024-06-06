package controllers

import "github.com/gin-gonic/gin"

type MemberController struct {
}

func (c MemberController) GetMmeber(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "asdf",
	})
}

func (c MemberController) CreateMember(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (c MemberController) DeleteMember(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
