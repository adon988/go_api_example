package controllers

import (
	"github.com/adon988/go_api_example/internal/middleware"
	models "github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/requests"
	"github.com/adon988/go_api_example/internal/responses"
	"github.com/adon988/go_api_example/internal/services"
	"github.com/adon988/go_api_example/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/plugin/dbresolver"
)

type AuthController struct {
	InfoDb utils.InfoDb
}

// @Summary Login
// @Description Login
// @Tags auth
// @Accept  json
// @Produce  json
// @param req body requests.LoginRequeset true "req"
// @Failure 400 {object} responses.ResponseFail "{"code": 100001, "msg":"", "data": {}}"
// @Failure 400 {object} responses.ResponseFail "{"code": 100002, "msg":"", "data": {}}"
// @success 200 {object} responses.LoginResonse    "{"code":0,"data":{"token":"token"},msg":"success"}"
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
		responses.FailWithErrorCode(responses.ACCOUNT_NOT_EXISTS, ctx)
		return
	}
	err := bcrypt.CompareHashAndPassword(auth.Password, []byte(req.Password))
	if err != nil {
		responses.FailWithErrorCode(responses.USERNAME_OR_PASSWORD_ERROR, ctx)
		return
	}

	token, _ := middleware.GenToken(auth.MemberId)

	data := responses.TokenResponse{
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

	// Use Write Mode: read user from sources `db1`
	tx := Db.Clauses(dbresolver.Write).Begin()

	authService := services.NewAuthService(tx)
	authType := "ApikeyAuth"
	auth := models.Authentication{
		MemberId: authId,
		Username: req.Username,
		Password: password,
		Type:     &authType,
	}

	err := authService.Register(auth)

	if err != nil {
		tx.Rollback()
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	tx.Commit()
	responses.Ok(ctx)
}
