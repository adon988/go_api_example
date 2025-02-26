package controllers

import (
	models "github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/requests"
	"github.com/adon988/go_api_example/internal/responses"
	"github.com/adon988/go_api_example/internal/services"
	"github.com/adon988/go_api_example/internal/utils"
	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	InfoDb utils.InfoDb
}

// @Summary Get Organization
// @Description Get all organizations that the member belongs to
// @Tags Organization
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} responses.OrganizationResponse
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /my/organization [get]
func (c OrganizationController) GetOrganizationMemberBelongTo(ctx *gin.Context) {
	Db, _ := c.InfoDb.InitDB()
	memberId := ctx.GetString("account")
	organizationService := services.NewOrganizationService(Db)
	var organizationsRes []responses.OrganizationResponse
	var err error
	organizations, err := organizationService.GetOrganizationMemberBelongTo(memberId)

	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	for _, organization := range organizations {
		organizationsRes = append(organizationsRes, responses.OrganizationResponse{
			Id:             organization.Id,
			Title:          organization.Title,
			Order:          organization.Order,
			SourceLanguage: organization.SourceLanguage,
			TargetLanguage: organization.TargetLanguage,
			Publish:        organization.Publish,
			CreaterId:      organization.CreaterId,
			CreatedAt:      organization.CreatedAt,
			UpdatedAt:      organization.UpdatedAt,
		})
	}

	responses.OkWithData(organizationsRes, ctx)
}

// @Summary Create Organization
// @Description Create a new organization, and the creator will be the admin of the organization
// @Tags Organization
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param organization body requests.OrganizationCreateRequest true "Organization object that needs to be created"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"success"}"
// @Failure 400 {string} string '{"code":-1, "data":{}, "msg":""}'
// @Router /my/organization [post]
func (c OrganizationController) CreateOrganization(ctx *gin.Context) {
	memberId := ctx.GetString("account")
	defaultRole := int32(1)
	var req requests.OrganizationCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	Db, _ := c.InfoDb.InitDB()
	organizationService := services.NewOrganizationService(Db)
	orgId, _ := utils.GenId()
	organizationData := models.Organization{
		Id:             orgId,
		Title:          req.Title,
		Order:          req.Order,
		SourceLanguage: req.SourceLanguage,
		TargetLanguage: req.TargetLanguage,
		Publish:        req.Publish,
		CreaterId:      memberId,
	}
	err := organizationService.CreateOrganizationNPermission(memberId, defaultRole, organizationData)

	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}

// @Summary Update Organization
// @Description Update organization information
// @Tags Organization
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param organization body requests.OrganizationUpdateRequest true "Organization object that needs to be updated"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"success"}"
// @Failure 400 {string} string '{"code":-1, "data":{}, "msg":""}'
// @Router /my/organization [patch]
func (c OrganizationController) UpdateOrganization(ctx *gin.Context) {
	var req requests.OrganizationUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	Db, _ := c.InfoDb.InitDB()
	organizationService := services.NewOrganizationService(Db)
	memberId := ctx.GetString("account")

	_, err := organizationService.IsMemberWithEditorPermissionOnOrganization(memberId, req.Id)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	organizationReq := models.Organization{
		Id:             req.Id,
		Title:          req.Title,
		Order:          req.Order,
		SourceLanguage: req.SourceLanguage,
		TargetLanguage: req.TargetLanguage,
		Publish:        req.Publish,
	}
	err = organizationService.UpdateOrganization(organizationReq)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}

// @Summary Delete Organization
// @Description Delete organization
// @Tags Organization
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param organization body requests.OrganizationDeleteRequest true "Organization object that needs to be deleted"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"success"}"
// @Failure 400 {string} string '{"code":-1, "data":{}, "msg":""}'
// @Router /my/organization [delete]
func (c OrganizationController) DeleteOrganization(ctx *gin.Context) {
	var req requests.OrganizationDeleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	Db, _ := c.InfoDb.InitDB()

	organizationService := services.NewOrganizationService(Db)
	memberId := ctx.GetString("account")

	_, err := organizationService.IsMemberWithEditorPermissionOnOrganization(memberId, req.Id)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	err = organizationService.DeleteOrganization(req.Id)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
	}
	responses.Ok(ctx)
}

// @Summary Assign Organization Permission
// @Description Assign organization permission to another member
// @Tags Organization
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param organization body requests.AssignRoleToMemberRequest true "Organization object that needs to be assigned"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"success"}"
// @Failure 400 {string} string '{"code":-1, "data":{}, "msg":""}'
// @Router /my/organization/assign [post]
func (c OrganizationController) AssignOrganizationPermission(ctx *gin.Context) {
	var req requests.AssignRoleToMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	Db, _ := c.InfoDb.InitDB()

	organizationService := services.NewOrganizationService(Db)
	memberId := ctx.GetString("account")

	orgPerm, err := organizationService.IsMemberWithEditorPermissionOnOrganization(memberId, req.OrganizationId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	//check if the member has enough permission to assign role to member
	if orgPerm.Role > req.RoleId {
		responses.FailWithMessage("permission denied (2)", ctx)
		return
	}

	organizationData := models.OrganizationPermission{
		MemberId: req.MemberId,
		EntityId: orgPerm.EntityId,
		Role:     req.RoleId,
	}

	memberServices := services.NewMemberService(Db)
	_, err = memberServices.GetMemberInfo(req.MemberId)

	if err != nil {
		responses.FailWithMessage("member not found", ctx)
		return
	}

	err = organizationService.AssignOrganizationPermission(organizationData)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}
