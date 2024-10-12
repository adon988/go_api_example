package services

import (
	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/repository"
	"gorm.io/gorm"
)

type CourseServiceInterface interface {
	CreateCourseNPermission(member_id string, role int32, course models.Course) error
	GetCourse(member_id string) ([]models.Course, error)
	UpdateCourse(course models.Course) error
	DeleteCourse(id string) error
	GetCourseByMemberIDAndCourseID(member_id string, course_id string) (models.Course, error)
	GetCoursePermissionByMemberIDAndCourseID(member_id string, course_id string) (models.CoursePermission, error)
	AssignCoursePermission(coursePermission models.CoursePermission) error
}

func NewCourseSerive(db *gorm.DB) CourseServiceInterface {
	return &CourseService{
		course:           repository.NewCourseRepository(db),
		coursePermission: repository.NewCoursePermissionRepository(db),
	}
}

type CourseService struct {
	course           repository.CourseRepository
	coursePermission repository.CoursePermissionRepository
}

func (service CourseService) CreateCourseNPermission(member_id string, role int32, course models.Course) error {
	coursePermission := models.CoursePermission{
		MemberId: member_id,
		EntityId: course.Id,
		Role:     role,
	}

	result := service.coursePermission.CreateCoursePermission(coursePermission)
	if result != nil {
		return result
	}

	err := service.course.CreateCourse(course)
	if err != nil {
		return err
	}
	return nil
}

func (service CourseService) GetCourse(member_id string) ([]models.Course, error) {
	return service.course.GetCourseByMemberID(member_id)
}

func (service CourseService) GetCourseByMemberIDAndCourseID(member_id string, course_id string) (models.Course, error) {
	return service.course.GetCourseByMemberIDAndCourseID(member_id, course_id)
}

func (service CourseService) UpdateCourse(course models.Course) error {
	return service.course.UpdateCourse(course)
}

func (service CourseService) DeleteCourse(id string) error {
	if err := service.coursePermission.DeleteCoursePermission(id); err != nil {
		return err
	}
	return service.course.DeleteCourse(id)
}

func (service CourseService) GetCoursePermissionByMemberIDAndCourseID(member_id string, course_id string) (models.CoursePermission, error) {
	return service.coursePermission.GetCoursePermissionByMemberIDAndCourseID(member_id, course_id)
}

func (service CourseService) AssignCoursePermission(coursePermission models.CoursePermission) error {
	return service.coursePermission.AssignCoursePermission(coursePermission)
}
