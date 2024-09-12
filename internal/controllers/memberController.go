package controllers

import (
	"strings"
	"time"

	model "github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/services"
	"github.com/adon988/go_api_example/internal/utils"
	"github.com/adon988/go_api_example/internal/utils/requests"
	"github.com/adon988/go_api_example/internal/utils/responses"

	"github.com/gin-gonic/gin"
	"gorm.io/plugin/dbresolver"
)

type MemberController struct {
	InfoDb utils.InfoDb
}

// GetMmeberById retrieves a member by ID.
// @Summary Get a member by ID
// @Description Get a member by ID
// @Tags member
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} responses.GetMemberResonse
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":"member not found"}'
// @Router /member [get]
func (c MemberController) GetMemberInfo(ctx *gin.Context) {

	memberId, _ := ctx.Get("account")
	Db, _ := c.InfoDb.InitDB()

	memberServices := services.NewMemberService(Db)
	members, err := memberServices.GetMemberInfo(memberId.(string))

	if err != nil {
		responses.FailWithMessage("member not found", ctx)
		return
	}

	data := responses.MemberinfoResponse{
		ID:        members.Id,
		Name:      responses.NullableString(members.Name),
		Birthday:  responses.NullableDate(members.Birthday),
		Gender:    responses.NullableInt(members.Gender),
		Email:     responses.NullableString(members.Email),
		CreatedAt: members.CreatedAt,
		UpdatedAt: members.UpdatedAt,
	}

	responses.OkWithData(data, ctx)
}

// UpdateMember updates a member.
// @Summary Update a member
// @Description Update a member
// @Tags member
// @Accept  json
// @Produce  json
// @Param req body requests.MemberUpdateRequest true "req" default({"name":"test2", "email":"","gender":0, "birthday":"2021-01-01"})
// @Security ApiKeyAuth
// @success 200 {object} responses.ResponseSuccess
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":"failed to update member"}'
// @Router /member [patch]
func (c MemberController) UpdateMember(ctx *gin.Context) {

	memberId, _ := ctx.Get("account")

	var req requests.MemberUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	Db, _ := c.InfoDb.InitDB()
	name := strings.TrimSpace(req.Name)
	birthday, _ := time.Parse("2006-01-02", req.Birthday) // Convert req.Birthday string to time.Time
	member := model.Member{
		Name:     &name,
		Birthday: &birthday,
		Email:    &req.Email,
		Gender:   &req.Gender,
	}

	memberService := services.NewMemberService(Db)
	result := memberService.UpdateMember(memberId.(string), member)

	if result != nil {
		responses.FailWithMessage("update member error", ctx)
		return
	}

	responses.Ok(ctx)
}

// DeleteMember deletes a member.
// @Summary Delete a member
// @Description Delete a member
// @Tags member
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @success 200 {object} responses.ResponseSuccess
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":"failed to delete member"}'
// @Router /member [delete]
func (c MemberController) DeleteMember(ctx *gin.Context) {

	memberId, _ := ctx.Get("account")

	Db, _ := c.InfoDb.InitDB()

	// Use Write Mode: read user from sources `db1`
	tx := Db.Clauses(dbresolver.Write).Begin()

	// Delete member and auth
	err := services.NewMemberService(Db).DeleteMemberAndAuth(memberId.(string))
	if err != nil {
		tx.Rollback()
		responses.FailWithMessage("failed to delete member", ctx)
		return
	}
	tx.Commit()
	responses.Ok(ctx)
}
