package controllers

import (
	"github.com/adon988/go_api_example/internal/utils"
	"github.com/adon988/go_api_example/internal/utils/responses"
	"github.com/gin-gonic/gin"
)

type CourseController struct {
	InfoDb utils.InfoDb
}

type CourseResponse struct {
	Id             string `json:"id"`
	Title          string `json:"title"`
	OrganizationId string `json:"organization_id"`
	Order          int32  `json:"order"`
	Publish        int32  `json:"publish"`
	CreaterId      string `json:"creator_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

func (c CourseController) GetCourse(ctx *gin.Context) {

	responses.Ok(ctx)
}

func (c CourseController) CreateCourse(ctx *gin.Context) {

	responses.Ok(ctx)
}

func (c CourseController) UpdateCourse(ctx *gin.Context) {

	responses.Ok(ctx)
}

func (c CourseController) DeleteCourse(ctx *gin.Context) {

	responses.Ok(ctx)
}
