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

func TestCourseRepository_GetCourseMmeberBelongTo(t *testing.T) {
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

func TestCourseRepository_GetCourseByMemberIDAndCourseID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Course{}, &models.CoursePermission{})
	repo := NewCourseRepository(mockDB)
	memberId := "1"
	unAuthMemberId := "2"
	entityId := "1"
	public := int32(1)
	unPublish := int32(0)
	course := models.Course{
		Id:        entityId,
		Title:     "course title",
		Order:     1,
		Publish:   public,
		CreaterId: memberId,
	}

	mockDB.Create(&course)
	course_perm := models.CoursePermission{
		MemberId: memberId,
		EntityId: entityId,
		Role:     1,
	}
	mockDB.Create(&course_perm)

	//check public scenario by auth member
	courseData, err := repo.GetCourseByMemberIDAndCourseID(memberId, entityId)
	assert.Nil(t, err)
	assert.Equal(t, course.Id, courseData.Id)
	assert.Equal(t, course.Title, courseData.Title)

	//check public scenario by un-auth member
	courseData, err = repo.GetCourseByMemberIDAndCourseID(unAuthMemberId, entityId)
	assert.Nil(t, err)
	assert.Equal(t, course.Id, courseData.Id)
	assert.Equal(t, course.Title, courseData.Title)

	// change public status to un-publish
	course.Publish = unPublish
	mockDB.Save(&course)

	// check private status by auth member
	courseData, err = repo.GetCourseByMemberIDAndCourseID(memberId, entityId)
	assert.Nil(t, err)
	assert.Equal(t, course.Id, courseData.Id)
	assert.Equal(t, course.Title, courseData.Title)

	//check non-public scenario by un-auth member
	courseData, err = repo.GetCourseByMemberIDAndCourseID(unAuthMemberId, entityId)
	assert.Error(t, err)
	assert.Equal(t, courseData, models.Course{})
}
