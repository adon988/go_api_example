package controllers

import (
	"fmt"

	"github.com/adon988/go_api_example/api/middleware"
	"github.com/adon988/go_api_example/api/response"
	"github.com/adon988/go_api_example/models"
	"github.com/adon988/go_api_example/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
}

type LoginVerify struct {
	Username string `json:"username" binding:"required" message:"username is required"`
	Password string `json:"password" binding:"required" message:"password is required"`
}

// @Summary Login
// @Description Login
// @Tags auth
// @Accept  json
// @Produce  json
// @param req body LoginVerify true "req"
// @error 400 {string} string "username not exists"
// @error 400 {string} string "username or password error"
// @Success 200 {string} string "ok"
// @Router /auth/login [post]
func (AuthController) Login(ctx *gin.Context) {
	var req LoginVerify
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	Db, _ := InfoDb.InitDB()
	var user models.Member
	result := Db.Where("username = ?", req.Username).First(&user)

	if result.RowsAffected == 0 {
		ctx.JSON(400, gin.H{"error": "username not exists"})
		return
	}
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "username or password error"})
		return
	}

	token, _ := middleware.GenToken(user.Id)

	data := gin.H{
		"token": token,
	}
	response.OkWithData(data, ctx)
}

// @Summary Register
// @Description Register
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @param req body LoginVerify true "req"
// @error 400 {string} string "username already exists"
// @error 400 {string} string "failed to create member"
// @Router /auth/register [post]
func (AuthController) Register(ctx *gin.Context) {
	var req LoginVerify
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	fmt.Printf("username: %s, password: %s, bcrypt: %s", req.Username, req.Password, password)

	Db, _ := InfoDb.InitDB()
	var user models.Member
	result := Db.First(&user, "username = ?", req.Username)

	if result.RowsAffected > 0 {
		ctx.JSON(400, gin.H{"error": "username already exists"})
		return
	}
	authId, err := utils.GenId()
	if err != nil {
		response.FailWithMessage("failed to create member", ctx)
		return
	}

	auth := models.Member{
		Id:       authId,
		Username: req.Username,
		Password: password,
	}

	result = Db.Create(&auth)

	if result.Error != nil {
		response.FailWithMessage("failed to create member", ctx)
		return
	}
	response.Ok(ctx)
}
