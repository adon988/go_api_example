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
	ID      string `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Age     int64  `json:"age" binding:"required, gt=0, lt=150"`
	Email   string `json:"email" binding:"email"`
	Address string `json:"address"`
}
type MemberCreateVerify struct {
	Name    string `form:"name" json:"name" binding:"required"`
	Age     int64  `form:"age" json:"age" binding:"required, gt=0, lt=15"`
	Email   string `form:"email" json:"email" binding:"email"`
	Address string `form:"address" json:"address"`
}

var InfoDb utils.InfoDb

func (c MemberController) GetMmeberById(ctx *gin.Context) {
	Db, _ := InfoDb.InitDB()

	memberID := ctx.Params.ByName("id")

	var members model.Member
	results := Db.First(&members, memberID)

	if results.Error != nil {
		response.FailWithMessage("member not found", ctx)
		return
	}
	data := gin.H{
		"member": gin.H{
			"ID":        members.ID,
			"Name":      members.Name,
			"Age":       members.Age,
			"Email":     members.Email,
			"CreatedAt": members.CreatedAt,
			"UpdatedAt": members.UpdatedAt,
		},
	}
	response.OkWithData(data, ctx)
}

func (c MemberController) UpdateMember(ctx *gin.Context) {

	var json MemberUpdateVerify
	if err := ctx.ShouldBindJSON(&json); err != nil {
		response.Fail(ctx)
		return
	}

	Db, _ := InfoDb.InitDB()

	member := model.Member{
		ID:      json.ID,
		Name:    strings.TrimSpace(json.Name),
		Age:     json.Age,
		Email:   &json.Email,
		Address: &json.Address,
	}

	result := Db.Where("id = ?", json.ID).Updates(&member)

	if result.Error != nil {
		response.FailWithMessage("update member error", ctx)
		return
	}

	response.Ok(ctx)
}

func (c MemberController) CreateMember(ctx *gin.Context) {
	var json MemberCreateVerify
	if err := ctx.ShouldBindJSON(&json); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	memberId, err := utils.GenId()
	if err != nil {
		response.FailWithMessage("failed to create member", ctx)
		return
	}

	member := model.Member{
		ID:      memberId,
		Name:    strings.TrimSpace(json.Name),
		Age:     json.Age,
		Email:   &json.Email,
		Address: &json.Address,
	}

	Db, _ := InfoDb.InitDB()
	result := Db.Create(&member)

	if result.Error != nil {
		response.FailWithMessage("failed to create member", ctx)
		return
	}
	response.Ok(ctx)
}

func (c MemberController) DeleteMember(ctx *gin.Context) {
	Db, _ := InfoDb.InitDB()
	member_id := ctx.PostForm("id")

	result := Db.Where("id = ?", member_id).Delete(&model.Member{})
	if result.Error != nil {
		response.FailWithMessage("failed to delete member", ctx)
		return
	}
	response.Ok(ctx)
}
