package repository

import (
	"fmt"

	"github.com/adon988/go_api_example/internal/models"
	"gorm.io/gorm"
)

type CoursePermissionRepository interface {
	CreateCoursePermission(coursePermission models.CoursePermission) error
	GetCoursePermissionByMemberID(member_id string) ([]models.CoursePermission, error)
	GetCoursePermissionByMemberIDAndCourseID(member_id string, course_id string) (models.CoursePermission, error)
	UpdateCoursePermission(coursePermission models.CoursePermission) error
	DeleteCoursePermission(id string) error
	AssignCoursePermission(coursePermission models.CoursePermission) error
}

func NewCoursePermissionRepository(db *gorm.DB) CoursePermissionRepository {
	return &CoursePermissionImpl{DB: db}
}

type CoursePermissionImpl struct {
	DB *gorm.DB
}

func (r *CoursePermissionImpl) CreateCoursePermission(coursePermission models.CoursePermission) error {
	hasRole := false
	for _, v := range models.PermissionRoleEnum {
		if v == coursePermission.Role {
			hasRole = true
			break
		}
	}
	if !hasRole {
		return fmt.Errorf("role not found")
	}
	result := r.DB.Create(&coursePermission)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CoursePermissionImpl) GetCoursePermissionByMemberID(member_id string) ([]models.CoursePermission, error) {
	var coursesPerms []models.CoursePermission
	result := r.DB.Find(&coursesPerms, "member_id = ?", member_id)
	if result.Error != nil {
		return nil, result.Error
	}
	return coursesPerms, nil
}

func (r *CoursePermissionImpl) GetCoursePermissionByMemberIDAndCourseID(member_id string, course_id string) (models.CoursePermission, error) {
	var coursePerms models.CoursePermission
	result := r.DB.Find(&coursePerms, "member_id = ? AND entity_id = ?", member_id, course_id)
	if result.Error != nil {
		return coursePerms, result.Error
	}
	if result.RowsAffected == 0 {
		return coursePerms, gorm.ErrRecordNotFound
	}
	return coursePerms, nil
}

func (r *CoursePermissionImpl) UpdateCoursePermission(course_permission models.CoursePermission) error {
	result := r.DB.Updates(&course_permission)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CoursePermissionImpl) DeleteCoursePermission(id string) error {
	result := r.DB.Delete(&models.CoursePermission{}, "entity_id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *CoursePermissionImpl) AssignCoursePermission(coursePermission models.CoursePermission) error {
	hasRole := false
	for _, v := range models.PermissionRoleEnum {
		if v == coursePermission.Role {
			hasRole = true
			break
		}
	}
	if !hasRole {
		return fmt.Errorf("role not found")
	}

	result := r.DB.Model(&models.CoursePermission{}).Where("member_id = ? AND entity_id = ?", coursePermission.MemberId, coursePermission.EntityId).Updates(&coursePermission)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		result = r.DB.Create(&coursePermission)
	}
	return result.Error
}
