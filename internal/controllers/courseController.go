package controllers

import (
	models "github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/services"
	"github.com/adon988/go_api_example/internal/utils"
	"github.com/adon988/go_api_example/internal/utils/requests"
	"github.com/adon988/go_api_example/internal/utils/responses"
	"github.com/gin-gonic/gin"
)

type CourseController struct {
	InfoDb utils.InfoDb
}

// @Summary Get Course
// @Description Get all courses that the member belongs to
// @Tags Course
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} responses.CourseResponse
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /admin/course [get]
func (c CourseController) GetCourse(ctx *gin.Context) {
	Db, _ := c.InfoDb.InitDB()
	memberId, _ := ctx.Get("account")
	courseService := services.NewCourseSerive(Db)
	var coursesRes []responses.CourseResponse
	var err error
	courses, err := courseService.GetCourse(memberId.(string))

	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	for _, course := range courses {
		coursesRes = append(coursesRes, responses.CourseResponse{
			Id:             course.Id,
			Title:          course.Title,
			OrganizationId: course.OrganizationId,
			Order:          course.Order,
			Publish:        course.Publish,
			CreaterId:      course.CreaterId,
			CreatedAt:      course.CreatedAt,
			UpdatedAt:      course.UpdatedAt,
		})
	}

	responses.OkWithData(coursesRes, ctx)
}

// @Summary Create Course
// @Description Create a course, and assign the creator as the admin of the course
// @Tags Course
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param title body requests.CourseCreateRequest true "Course object that needs to be created"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /admin/course [post]
func (c CourseController) CreateCourse(ctx *gin.Context) {
	memberId, _ := ctx.Get("account")
	defaultRole := int32(1)
	var req requests.CourseCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	Db, _ := c.InfoDb.InitDB()

	organizationService := services.NewOrganizationService(Db)
	_, err := organizationService.IsMemberWithEditorPermissionOnOrganization(memberId.(string), req.OrganizationId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	courseService := services.NewCourseSerive(Db)
	courseId, _ := utils.GenId()
	courseData := models.Course{
		Id:             courseId,
		Title:          req.Title,
		OrganizationId: req.OrganizationId,
		Order:          req.Order,
		Publish:        req.Publish,
		CreaterId:      memberId.(string),
	}
	err = courseService.CreateCourseNPermission(memberId.(string), defaultRole, courseData)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}

// @Summary Update Course
// @Description Update course information
// @Tags Course
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param course body requests.CourseUpdateReqeust true "Course object that needs to be updated"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /admin/course [put]
func (c CourseController) UpdateCourse(ctx *gin.Context) {
	var req requests.CourseUpdateReqeust
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	Db, _ := c.InfoDb.InitDB()
	courseService := services.NewCourseSerive(Db)
	memberId := ctx.GetString("account")

	organizationService := services.NewOrganizationService(Db)
	_, err := organizationService.IsMemberWithEditorPermissionOnOrganization(memberId, req.OrganizationId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	_, err = courseService.IsMemberWithEditorPermissionOnCourse(memberId, req.Id)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	coursere := models.Course{
		Id:             req.Id,
		Title:          req.Title,
		OrganizationId: req.OrganizationId,
		Order:          req.Order,
		Publish:        req.Publish,
	}

	err = courseService.UpdateCourse(coursere)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}

// @Summary Delete Course
// @Description Delete course
// @Tags Course
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param course body requests.CourseDeleteReqeust true "Course object that needs to be deleted"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /admin/course [delete]
func (c CourseController) DeleteCourse(ctx *gin.Context) {
	var req requests.CourseDeleteReqeust
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	Db, _ := c.InfoDb.InitDB()

	courseService := services.NewCourseSerive(Db)
	memberId := ctx.GetString("account")

	_, err := courseService.IsMemberWithEditorPermissionOnCourse(memberId, req.Id)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	err = courseService.DeleteCourse(req.Id)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	responses.Ok(ctx)
}

// @Summary Assign Course Permission
// @Description Assign course permission to member
// @Tags Course
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param course body requests.AssignCourseRequest true "Course object that needs to be assigned"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /admin/course/assign [post]
func (c CourseController) AssignCoursePermission(ctx *gin.Context) {
	var req requests.AssignCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	Db, _ := c.InfoDb.InitDB()
	courseService := services.NewCourseSerive(Db)
	memberId := ctx.GetString("account")

	coursePerm, err := courseService.IsMemberWithEditorPermissionOnCourse(memberId, req.CourseId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	//check if the member has enough permission to assign role to member
	if coursePerm.Role > req.RoleId {
		responses.FailWithMessage("permission denied (2)", ctx)
		return
	}

	courseData := models.CoursePermission{
		MemberId: req.MemberId,
		EntityId: coursePerm.EntityId,
		Role:     req.RoleId,
	}
	memberServices := services.NewMemberService(Db)
	_, err = memberServices.GetMemberInfo(req.MemberId)

	if err != nil {
		responses.FailWithMessage("member not found", ctx)
		return
	}

	err = courseService.AssignCoursePermission(courseData)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}
