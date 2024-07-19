package controllers

import (
	"github.com/adon988/go_api_example/api/services"
	"github.com/adon988/go_api_example/models"
	"github.com/adon988/go_api_example/models/requests"
	"github.com/adon988/go_api_example/models/responses"
	"github.com/adon988/go_api_example/utils"
	"github.com/gin-gonic/gin"
)

type RoleResponse struct {
	Id       int32  `json:"id"`
	Title    string `json:"title"`
	RoleType string `json:"role_type"`
	Image    string `json:"image"`
}

type RoleController struct {
	InfoDb utils.InfoDb
}

// @Summary Get Role
// @Description Get all roles
// @Tags Role
// @Accept json
// @Produce json
// @Success 200 {array} RoleResponse
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /roles [get]
func (c RoleController) GetRole(ctx *gin.Context) {
	Db, _ := c.InfoDb.InitDB()
	roleService := services.NewRoleService(Db)

	roles, err := roleService.GetRole()

	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	var roleResponses []RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, RoleResponse{
			Id:       role.Id,
			Title:    role.Title,
			RoleType: role.RoleType,
			Image:    role.Image,
		})
	}

	responses.OkWithData(roleResponses, ctx)
}

// @Summary Create Role
// @Description Create a new role
// @Tags Role
// @Accept json
// @Produce json
// @Param role body requests.RoleCreateRequest true "Role object that needs to be created"
// @Success 200 {string} string "Role created successfully"
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /roles [post]
func (c RoleController) CreateRole(ctx *gin.Context) {

	var req requests.RoleCreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	Db, _ := c.InfoDb.InitDB()
	roleService := services.NewRoleService(Db)

	role := models.Role{
		Title:    req.Title,
		RoleType: req.RoleType,
		Image:    req.Image,
	}

	err := roleService.CreateRole(role)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	responses.OkWithMessage("Role created successfully", ctx)
}

// @Summary Update Role
// @Description Update an existing role
// @Tags Role
// @Accept json
// @Produce json
// @Param role body requests.RoleUpdateRequest true "Role object that needs to be updated"
// @Success 200 {string} string "Role updated successfully"
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /roles [put]
func (c RoleController) UpdateRole(ctx *gin.Context) {
	var req requests.RoleUpdateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	Db, _ := c.InfoDb.InitDB()
	roleService := services.NewRoleService(Db)

	role := models.Role{
		Id:       req.Id,
		Title:    req.Title,
		RoleType: req.RoleType,
		Image:    req.Image,
	}

	err := roleService.UpdateRole(role)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	responses.OkWithMessage("Role updated successfully", ctx)
}

// @Summary Delete Role
// @Description Delete an existing role
// @Tags Role
// @Accept json
// @Produce json
// @Param role body requests.RoleDeleteRequest true "Role object that needs to be deleted"
// @Success 200 {string} string "Role deleted successfully"
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /roles [delete]
func (c RoleController) DeleteRole(ctx *gin.Context) {
	var req requests.RoleDeleteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	Db, _ := c.InfoDb.InitDB()
	roleService := services.NewRoleService(Db)

	err := roleService.DeleteRole(req.Id)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	responses.OkWithMessage("Role deleted successfully", ctx)
}
