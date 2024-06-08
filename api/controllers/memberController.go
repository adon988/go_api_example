package controllers

import (
	"strings"

	"github.com/adon988/go_api_example/api/response"
	model "github.com/adon988/go_api_example/models"
	"github.com/adon988/go_api_example/utils"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

type MemberUpdateVerify struct {
	Name    string `json:"name" binding:""`
	Age     int64  `json:"age" binding:""`
	Email   string `json:"email" binding:"email"`
	Address string `json:"address"`
}

type MemberCreateVerify struct {
	Name    string `form:"name" json:"name" binding:""`
	Age     int64  `form:"age" json:"age" binding:""`
	Email   string `form:"email" json:"email" binding:"email"`
	Address string `form:"address" json:"address"`
}

var InfoDb utils.InfoDb

// GetMmeberById retrieves a member by ID.
// @Summary Get a member by ID
// @Description Get a member by ID
// @Tags member
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} MemberController
// @Failure 400 {string} string "member not found"
// @Router /member [get]
func (c MemberController) GetMmeberInfo(ctx *gin.Context) {

	memberId, _ := ctx.Get("account")
	Db, _ := InfoDb.InitDB()

	var members model.Member
	results := Db.First(&members, memberId)

	if results.Error != nil {
		response.FailWithMessage("member not found", ctx)
		return
	}
	data := gin.H{
		"member": gin.H{
			"ID":        members.Id,
			"Name":      members.Name,
			"Age":       members.Age,
			"Email":     members.Email,
			"CreatedAt": members.CreatedAt,
			"UpdatedAt": members.UpdatedAt,
		},
	}
	response.OkWithData(data, ctx)
}

// UpdateMember updates a member.
// @Summary Update a member
// @Description Update a member
// @Tags member
// @Accept  json
// @Produce  json
// @Param req body MemberUpdateVerify true "req" default({"name":"test","age":18,"email":"","address":""})
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "failed to update member"
// @Router /member [patch]
func (c MemberController) UpdateMember(ctx *gin.Context) {

	memberId, _ := ctx.Get("account")

	var req MemberUpdateVerify
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Fail(ctx)
		return
	}

	Db, _ := InfoDb.InitDB()
	name := strings.TrimSpace(req.Name)
	member := model.Member{
		Name:    &name,
		Age:     &req.Age,
		Email:   &req.Email,
		Address: &req.Address,
	}

	result := Db.Where("id = ?", memberId).Updates(&member)

	if result.Error != nil {
		response.FailWithMessage("update member error", ctx)
		return
	}

	response.Ok(ctx)
}

// DeleteMember deletes a member.
// @Summary Delete a member
// @Description Delete a member
// @Tags member
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "failed to delete member"
// @Router /member [delete]
func (c MemberController) DeleteMember(ctx *gin.Context) {
	Db, _ := InfoDb.InitDB()
	memberId, _ := ctx.Get("account")

	result := Db.Where("id = ?", memberId).Delete(&model.Member{})
	if result.Error != nil {
		response.FailWithMessage("failed to delete member", ctx)
		return
	}
	response.Ok(ctx)
}
