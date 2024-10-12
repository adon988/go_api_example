package services

import (
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCourseService_CreateCourseNPermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Course{}, &models.CoursePermission{})

	memberId := "1"
	course := models.Course{
		Id:        "1",
		Title:     "course title",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}
	role := int32(1)
	service := NewCourseSerive(mockDB)
	result := service.CreateCourseNPermission(memberId, role, course)
	assert.Nil(t, result)

}

func TestCourseService_GetCourse(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Course{}, &models.CoursePermission{})

	memberId := "1"
	course := models.Course{
		Id:        "1",
		Title:     "course title",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}
	role := int32(1)
	service := NewCourseSerive(mockDB)
	result := service.CreateCourseNPermission(memberId, role, course)
	assert.Nil(t, result)

	courses, err := service.GetCourse(memberId)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(courses))
}

func TestCourseService_GetCoursePermissionByMemberIDAndCourseID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Course{}, &models.CoursePermission{})

	memberId := "1"
	course := models.Course{
		Id:        "1",
		Title:     "course title",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}
	role := int32(1)
	service := NewCourseSerive(mockDB)
	result := service.CreateCourseNPermission(memberId, role, course)
	assert.Nil(t, result)

	coursePerms, err := service.GetCoursePermissionByMemberIDAndCourseID(course.Id, memberId)
	assert.Nil(t, err)
	assert.Equal(t, "1", coursePerms.EntityId)
}

func TestCourseService_UpdateCourse(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Course{}, &models.CoursePermission{})

	memberId := "1"
	course := models.Course{
		Id:        "1",
		Title:     "course title",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}
	role := int32(1)
	service := NewCourseSerive(mockDB)
	result := service.CreateCourseNPermission(memberId, role, course)
	assert.Nil(t, result)

	course.Title = "new course title"
	err := service.UpdateCourse(course)
	assert.Nil(t, err)

	courses, err := service.GetCourse(memberId)
	assert.Nil(t, err)
	assert.Equal(t, "new course title", courses[0].Title)
}

func TestCourseService_DeleteCourse(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Course{}, &models.CoursePermission{})

	memberId := "1"
	course := models.Course{
		Id:        "1",
		Title:     "course title",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}
	role := int32(1)
	service := NewCourseSerive(mockDB)
	result := service.CreateCourseNPermission(memberId, role, course)
	assert.Nil(t, result)

	err := service.DeleteCourse(course.Id)
	assert.Nil(t, err)

	courses, err := service.GetCourse(memberId)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(courses))
}

func TestCourseService_GetCourseByMemberIDAndCourseID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Course{}, &models.CoursePermission{})
	repo := NewCourseSerive(mockDB)

	course := models.Course{
		Id:        "1",
		Title:     "course title",
		Order:     1,
		Publish:   1,
		CreaterId: "1",
	}

	mockDB.Create(&course)
	course_perm := models.CoursePermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	mockDB.Create(&course_perm)
	courseData, err := repo.GetCourseByMemberIDAndCourseID("1", "1")
	assert.Nil(t, err)
	assert.Equal(t, course.Id, courseData.Id)
	assert.Equal(t, course.Title, courseData.Title)
}
