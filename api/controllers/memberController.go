package controllers

import (
	"strings"
	"time"

	"github.com/adon988/go_api_example/api/services"
	model "github.com/adon988/go_api_example/models"
	"github.com/adon988/go_api_example/models/requests"
	"github.com/adon988/go_api_example/models/responses"
	"github.com/adon988/go_api_example/utils"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
	InfoDb utils.InfoDb
}

var InfoDb utils.InfoDb

type MemberinfoResponse struct {
	ID        string    `json:"id" example:"123456"`
	Name      string    `json:"name" example:"test"`
	Birthday  string    `json:"birthday" example:"2021-01-01"`
	Email     string    `json:"email" example:"example@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2021-01-01 00:00:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-01-01 00:00:00"`
}
type GetMemberResonse struct {
	Code int `json:"code" example:"0"`
	Data MemberinfoResponse
	Msg  string `json:"msg" example:"success"`
}

// GetMmeberById retrieves a member by ID.
// @Summary Get a member by ID
// @Description Get a member by ID
// @Tags member
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} GetMemberResonse
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

	data := MemberinfoResponse{
		ID:        members.Id,
		Name:      responses.NullableString(members.Name),
		Birthday:  responses.NullableDate(members.Birthday),
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
// @Param req body requests.MemberUpdateRequesettrue "req" default({"name":"test","age":18,"email":"","address":""})
// @Security ApiKeyAuth
// @success 200 {object} responses.ResponseSuccess
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":"failed to update member"}'
// @Router /member [patch]
func (c MemberController) UpdateMember(ctx *gin.Context) {

	memberId, _ := ctx.Get("account")

	var req requests.MemberUpdateRequeset
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

	result := Db.Where("id = ?", memberId).Updates(&member)

	if result.Error != nil {
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
	tx := Db.Begin() // start a transaction

	deleteMemberResult := tx.Where("id = ?", memberId).Delete(&model.Member{})
	if deleteMemberResult.Error != nil {
		tx.Rollback()
		responses.FailWithMessage("failed to delete member (:0)", ctx)
		return
	}
	deleteAuthResult := tx.Where("member_id = ?", memberId).Delete(&model.Authentication{})
	if deleteAuthResult.Error != nil {
		tx.Rollback()
		responses.FailWithMessage("failed to delete member (:1)", ctx)
		return
	}
	tx.Commit()
	responses.Ok(ctx)
}
