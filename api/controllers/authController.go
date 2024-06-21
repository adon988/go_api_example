package controllers

import (
	"github.com/adon988/go_api_example/api/middleware"
	models "github.com/adon988/go_api_example/models"
	"github.com/adon988/go_api_example/models/requests"
	"github.com/adon988/go_api_example/models/responses"
	"github.com/adon988/go_api_example/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	InfoDb utils.InfoDb
}

type TokenResponse struct {
	Token string `json:"token" example:"jwt token"`
}
type LoginResonse struct {
	Code int `json:"code" example:"0"`
	Data TokenResponse
	Msg  string `json:"msg" example:"success"`
}

// @Summary Login
// @Description Login
// @Tags auth
// @Accept  json
// @Produce  json
// @param req body LoginVerify true "req"
// @Failure 400 {object} responses.ResponseFail "msg: username or password error"
// @Failure 400 {object} responses.ResponseFail "msg: account not exists"
// @success 200 {object} LoginResonse    "{"code":0,"data":{"token":"token"},msg":"success"}"
// @Router /auth/login [post]
func (c AuthController) Login(ctx *gin.Context) {
	var req requests.LoginRequeset
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
	}

	Db, _ := c.InfoDb.InitDB()
	var auth models.Authentication
	result := Db.Where("username = ?", req.Username).First(&auth)

	if result.RowsAffected == 0 {
		responses.FailWithMessage("account not exists", ctx)
		return
	}
	err := bcrypt.CompareHashAndPassword(auth.Password, []byte(req.Password))
	if err != nil {
		responses.FailWithMessage("username or password error", ctx)
		return
	}

	token, _ := middleware.GenToken(auth.MemberId)

	data := TokenResponse{
		Token: token,
	}
	responses.OkWithData(data, ctx)
}

// @Summary Register
// @Description Register
// @Tags auth
// @Accept  json
// @Produce  json
// @success 200 {object} responses.ResponseSuccess
// @param req body requests.LoginRequeset true "req"
// @Failure 400 {object} responses.ResponseFail "msg: account already exists(:0) \n msg: failed to create account(:1, :2)"
// @Router /auth/register [post]
func (c AuthController) Register(ctx *gin.Context) {
	var req requests.LoginRequeset
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	authId, _ := utils.GenId()
	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	Db, _ := c.InfoDb.InitDB()

	tx := Db.Begin() // start a transaction

	var auth models.Authentication
	result := tx.First(&auth, "username = ?", req.Username)

	if result.RowsAffected > 0 {
		tx.Rollback()
		responses.FailWithMessage("account already exists(:0)", ctx)
		return
	}

	// generate account
	auth = models.Authentication{
		MemberId: authId,
		Username: req.Username,
		Password: password,
	}

	authCreateResult := tx.Create(&auth)

	if authCreateResult.Error != nil {
		tx.Rollback()
		responses.FailWithMessage("failed to create account(:1)", ctx)
		return
	}

	// create member after account generate successfully
	member := models.Member{
		Id: authId,
	}
	memberCreateResult := tx.Create(&member)
	if memberCreateResult.Error != nil {
		tx.Rollback()
		responses.FailWithMessage("failed to create account (:2)", ctx)
		return
	}

	tx.Commit()
	responses.Ok(ctx)
}
