package controllers

import (
	models "github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/requests"
	"github.com/adon988/go_api_example/internal/responses"
	"github.com/adon988/go_api_example/internal/services"
	"github.com/adon988/go_api_example/internal/utils"
	"github.com/gin-gonic/gin"
)

type UnitController struct {
	InfoDb utils.InfoDb
}

// @Summary Get Units
// @Description Get all units that the member belongs to
// @Tags Unit
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param account header string true "Account"
// @Success 200 {array} responses.UnitResponse
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /my/unit [get]
func (c UnitController) GetUnitMemberBelongTo(ctx *gin.Context) {
	Db, _ := c.InfoDb.InitDB()
	memberId := ctx.GetString("account")
	unitService := services.NewUnitService(Db)
	var unitsRes []responses.UnitResponse
	var err error
	units, err := unitService.GetUnitMemberBelongTo(memberId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	for _, unit := range units {
		unitsRes = append(unitsRes, responses.UnitResponse{
			Id:        unit.Id,
			Title:     unit.Title,
			CourseId:  unit.CourseId,
			Order:     unit.Order,
			Publish:   unit.Publish,
			CreaterId: unit.CreaterId,
			CreatedAt: unit.CreatedAt,
			UpdatedAt: unit.UpdatedAt,
		})
	}

	responses.OkWithData(unitsRes, ctx)
}

// @Summary Create Unit
// @Description Create a unit
// @Tags Unit
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param title body requests.UnitCreateRequest true "Unit object that needs to be created"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /my/unit [post]
func (c UnitController) CreateUnit(ctx *gin.Context) {
	memberId := ctx.GetString("account")
	defaultRole := int32(1)
	var req requests.UnitCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	Db, _ := c.InfoDb.InitDB()

	courseServicde := services.NewCourseSerive(Db)
	_, err := courseServicde.IsMemberWithEditorPermissionOnCourse(ctx.GetString("account"), req.CourseId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	unitService := services.NewUnitService(Db)
	unitId, _ := utils.GenId()
	unitData := models.Unit{
		Id:        unitId,
		Title:     req.Title,
		CourseId:  req.CourseId,
		Order:     req.Order,
		Publish:   req.Publish,
		CreaterId: memberId,
	}
	err = unitService.CreateUnitNPermission(memberId, defaultRole, unitData)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}

// @Summary Update Unit
// @Description Update a unit
// @Tags Unit
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param unit body requests.UnitUpdateRequest true "Unit object that needs to be updated"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /my/unit [put]
func (c UnitController) UpdateUnit(ctx *gin.Context) {
	var req requests.UnitUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	Db, _ := c.InfoDb.InitDB()
	courseServicde := services.NewCourseSerive(Db)
	_, err := courseServicde.IsMemberWithEditorPermissionOnCourse(ctx.GetString("account"), req.CourseId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	unitService := services.NewUnitService(Db)
	memberId := ctx.GetString("account")
	_, err = unitService.IsMemberWithEditorPermissionOnUnit(memberId, req.Id)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	unit := models.Unit{
		Id:        req.Id,
		Title:     req.Title,
		CourseId:  req.CourseId,
		Order:     req.Order,
		Publish:   req.Publish,
		CreaterId: memberId,
	}
	err = unitService.UpdateUnit(unit)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	responses.Ok(ctx)
}

// @Summary Delete Unit
// @Description Delete a unit
// @Tags Unit
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param unit body requests.UnitDeleteRequest true "Unit object that needs to be deleted"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /my/unit [delete]
func (c UnitController) DeleteUnit(ctx *gin.Context) {
	var req requests.UnitDeleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	Db, _ := c.InfoDb.InitDB()
	unitService := services.NewUnitService(Db)
	memberId := ctx.GetString("account")
	_, err := unitService.IsMemberWithEditorPermissionOnUnit(memberId, req.Id)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	err = unitService.DeleteUnit(req.Id)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	responses.Ok(ctx)
}

// @Summary Assign Unit Permission
// @Description Assign a unit permission
// @Tags Unit
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param unit body requests.AssignUnitPermissionRequest true "Unit object that needs to be assigned"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /my/unit/permission [post]
func (c UnitController) AssignUnitPermission(ctx *gin.Context) {
	var req requests.AssignUnitPermissionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	Db, _ := c.InfoDb.InitDB()

	unitService := services.NewUnitService(Db)
	memberId := ctx.GetString("account")
	unitPerm, err := unitService.IsMemberWithEditorPermissionOnUnit(memberId, req.UnitId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	if unitPerm.Role > req.RoleId {
		responses.FailWithMessage("permission denied (2)", ctx)
		return
	}
	unitData := models.UnitPermission{
		MemberId: memberId,
		EntityId: unitPerm.EntityId,
		Role:     req.RoleId,
	}
	memberService := services.NewMemberService(Db)
	_, err = memberService.GetMemberInfo(req.MemberId)
	if err != nil {
		responses.FailWithMessage("member not found", ctx)
		return
	}
	err = unitService.AssignUnitPermission(unitData)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}

// @Summary Get Units By Course ID
// @Description Get all units by course id
// @Tags Unit
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param account header string true "Account"
// @Param course_id query string true "Course ID"
// @Success 200 {array} responses.UnitResponse
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /my/course/:course_id/units [get]
func (c UnitController) GetUnitsByCourseID(ctx *gin.Context) {
	memberId := ctx.GetString("account")
	courseId := ctx.Param("course_id")
	Db, _ := c.InfoDb.InitDB()
	courseService := services.NewCourseSerive(Db)
	//check Unit Permission
	_, err := courseService.GetCoursePermissionByMemberIDAndCourseID(memberId, courseId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	unitService := services.NewUnitService(Db)
	units, err := unitService.GetUnitsByCourseID(courseId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	var unitsRes []responses.UnitResponse
	for _, unit := range units {
		unitsRes = append(unitsRes, responses.UnitResponse{
			Id:        unit.Id,
			Title:     unit.Title,
			CourseId:  unit.CourseId,
			Order:     unit.Order,
			Publish:   unit.Publish,
			CreaterId: unit.CreaterId,
			CreatedAt: unit.CreatedAt,
			UpdatedAt: unit.UpdatedAt,
		})
	}
	responses.OkWithData(unitsRes, ctx)
}
