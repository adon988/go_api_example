package repository

import (
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCourseRepository_CreateCourse(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Course{}, &models.CoursePermission{})
	repo := NewCourseRepository(mockDB)
	memberId := "1"

	course := models.Course{
		Id:        "1",
		Title:     "course title",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}

	err := repo.CreateCourse(course)
	assert.Nil(t, err)
}

func TestCourseRepository_UpdateCourse(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Course{}, &models.CoursePermission{})
	repo := NewCourseRepository(mockDB)
	memberId := "1"

	course := models.Course{
		Id:        "1",
		Title:     "course title",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}

	err := mockDB.Create(&course)
	assert.Nil(t, err.Error)

	course.Title = "updated course title"
	result := repo.UpdateCourse(course)
	assert.Nil(t, result)
	var courses []models.Course
	mockDB.Find(&courses)
	assert.Equal(t, 1, len(courses))
	assert.Equal(t, "updated course title", courses[0].Title)
}

func TestCourseRepository_DeleteCourse(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Course{}, &models.CoursePermission{})
	repo := NewCourseRepository(mockDB)
	memberId := "1"

	course := models.Course{
		Id:        "1",
		Title:     "course title",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}

	err := mockDB.Create(&course)
	assert.Nil(t, err.Error)

	result := repo.DeleteCourse(course.Id)
	assert.Nil(t, result)
	var courses []models.Course
	mockDB.Find(&courses)
	assert.Equal(t, 0, len(courses))
}

func TestCourseRepository_GetCourse(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Course{}, &models.CoursePermission{})
	repo := NewCourseRepository(mockDB)

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

	courses, err := repo.GetCourseByMemberID("1")
	if err != nil {
		t.Errorf("Error getting courses: %v", err)
		return
	}
	assert.Equal(t, 1, len(courses))
	assert.Equal(t, course.Title, courses[0].Title)
}
