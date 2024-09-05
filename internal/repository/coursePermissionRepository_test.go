package repository

import (
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCoursePermissionRepository_CreateCoursePermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.CoursePermission{})
	repo := NewCoursePermissionRepository(mockDB)
	coursePerm := models.CoursePermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	err := repo.CreateCoursePermission(coursePerm)
	assert.Nil(t, err)
	// 檢查 course_permissions 是否被建立
	var coursePerms []models.CoursePermission
	mockDB.Find(&coursePerms)
	assert.Equal(t, 1, len(coursePerms))
	assert.Equal(t, "1", coursePerms[0].MemberId)
	assert.Equal(t, "1", coursePerms[0].EntityId)
	// test role not found
	coursePerm = models.CoursePermission{
		MemberId: "1",
		EntityId: "1",
		Role:     4,
	}
	err = repo.CreateCoursePermission(coursePerm)
	assert.Equal(t, "role not found", err.Error())

}

func TestCoursePermissionRepository_GetCourseByMemberIDAndCourseID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.CoursePermission{})
	repo := NewCoursePermissionRepository(mockDB)
	coursePerm := models.CoursePermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	//creatre
	err := repo.AssignCoursePermission(coursePerm)
	assert.Nil(t, err)
	coursePerms, err := repo.GetCoursePermissionByMemberIDAndCourseID("1", "1")
	assert.Nil(t, err)
	assert.Equal(t, "1", coursePerms.MemberId)
	assert.Equal(t, "1", coursePerms.EntityId)
	assert.Equal(t, int32(1), coursePerm.Role)
	//update
	coursePerm2 := models.CoursePermission{
		MemberId: "1",
		EntityId: "1",
		Role:     2,
	}
	err = repo.AssignCoursePermission(coursePerm2)
	assert.Nil(t, err)
	coursePerms, err = repo.GetCoursePermissionByMemberIDAndCourseID("1", "1")
	assert.Nil(t, err)
	assert.Equal(t, int32(2), coursePerm2.Role)
}

func TestCoursePermissionRepository_UpdateCoursePermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.CoursePermission{})
	repo := NewCoursePermissionRepository(mockDB)
	coursePerm := models.CoursePermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	mockDB.Create(&coursePerm)
	coursePerm.Role = 2
	err := repo.UpdateCoursePermission(coursePerm)
	assert.Nil(t, err)
	// 檢查 course_permissions 是否被更新
	var coursePerms []models.CoursePermission
	mockDB.Find(&coursePerms)
	assert.Equal(t, 1, len(coursePerms))
	assert.Equal(t, int32(2), coursePerms[0].Role)
}

func TestCoursePermissionRepository_DeleteCoursePermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.CoursePermission{})
	repo := NewCoursePermissionRepository(mockDB)
	coursePerm := models.CoursePermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	mockDB.Create(&coursePerm)
	err := repo.DeleteCoursePermission("1")
	assert.Nil(t, err)
	// 檢查 course_permissions 是否被刪除
	var coursePerms []models.CoursePermission
	mockDB.Find(&coursePerms)
	assert.Equal(t, 0, len(coursePerms))
}

func TestCoursePermissionRepository_AssignCoursePermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.CoursePermission{})
	repo := NewCoursePermissionRepository(mockDB)
	coursePerm := models.CoursePermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	//create
	err := repo.AssignCoursePermission(coursePerm)
	assert.Nil(t, err)
	// 檢查 course_permissions 是否被建立
	var coursePerms []models.CoursePermission
	mockDB.Find(&coursePerms)
	assert.Equal(t, 1, len(coursePerms))
	assert.Equal(t, "1", coursePerms[0].MemberId)
	assert.Equal(t, "1", coursePerms[0].EntityId)
	// test role not found
	coursePerm = models.CoursePermission{
		MemberId: "1",
		EntityId: "1",
		Role:     4,
	}
	err = repo.AssignCoursePermission(coursePerm)
	assert.Equal(t, "role not found", err.Error())
	//update
	coursePerm2 := models.CoursePermission{
		MemberId: "1",
		EntityId: "1",
		Role:     2,
	}
	err = repo.AssignCoursePermission(coursePerm2)
	assert.Nil(t, err)
	// 檢查 course_permissions 是否被更新
	mockDB.Find(&coursePerms)
	assert.Equal(t, 1, len(coursePerms))
	assert.Equal(t, int32(2), coursePerms[0].Role)
}
