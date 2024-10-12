package repository

import (
	"fmt"

	"github.com/adon988/go_api_example/internal/models"
	"gorm.io/gorm"
)

type CourseRepository interface {
	CreateCourse(course models.Course) error
	UpdateCourse(course models.Course) error
	DeleteCourse(id string) error
	GetCourseByMemberID(member_id string) ([]models.Course, error)
	GetCourseByMemberIDAndCourseID(member_id string, course_id string) (models.Course, error)
}

type CourseRepositoryImpl struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &CourseRepositoryImpl{DB: db}
}

func (r *CourseRepositoryImpl) CreateCourse(course models.Course) error {
	err := r.DB.Create(&course)

	if err != nil {
		return err.Error
	}

	return nil
}

func (r *CourseRepositoryImpl) UpdateCourse(course models.Course) error {
	err := r.DB.Where("id = ?", course.Id).Updates(course)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return fmt.Errorf("no course found with id: %s", course.Id)
	}
	return nil
}

func (r *CourseRepositoryImpl) DeleteCourse(id string) error {
	result := r.DB.Delete(&models.Course{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no course found with id: %s", id)
	}

	return nil
}

func (r *CourseRepositoryImpl) GetCourseByMemberID(member_id string) ([]models.Course, error) {
	var courses []models.Course
	err := r.DB.Model(&models.Course{}).Joins("JOIN course_permissions ON courses.id = course_permissions.entity_id").Where("course_permissions.member_id = ?", member_id).Find(&courses)

	if err.Error != nil {
		return nil, err.Error
	}
	return courses, nil
}

func (r *CourseRepositoryImpl) GetCourseByMemberIDAndCourseID(member_id string, course_id string) (models.Course, error) {
	var course models.Course
	err := r.DB.Model(&models.Course{}).Joins("JOIN course_permissions ON courses.id = course_permissions.entity_id").Where("course_permissions.member_id = ? AND courses.id = ?", member_id, course_id).Find(&course)
	if err.Error != nil {
		return models.Course{}, err.Error
	}
	return course, nil
}
