package services

import (
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUnitSerivce_CreateUnitNPermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Unit{}, &models.UnitPermission{})

	memberId := "1"
	unit := models.Unit{
		Id:        "1",
		Title:     "unit title",
		CourseId:  "1",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}
	role := int32(1)
	service := NewUnitService(mockDB)
	result := service.CreateUnitNPermission(memberId, role, unit)
	assert.Nil(t, result)
}

func TestUnitSerivce_GetUnitMemberBelongTo(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Unit{}, &models.UnitPermission{})

	memberId := "1"
	unit := models.Unit{
		Id:        "1",
		Title:     "unit title",
		CourseId:  "1",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}
	role := int32(1)
	service := NewUnitService(mockDB)
	result := service.CreateUnitNPermission(memberId, role, unit)
	assert.Nil(t, result)

	units, err := service.GetUnitMemberBelongTo(memberId)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(units))
}

func TestUnitSerivce_GetUnitPermissionByMemberIDAndUnitID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Unit{}, &models.UnitPermission{})

	memberId := "1"
	unit := models.Unit{
		Id:        "1",
		Title:     "unit title",
		CourseId:  "1",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}
	role := int32(1)
	service := NewUnitService(mockDB)
	result := service.CreateUnitNPermission(memberId, role, unit)
	assert.Nil(t, result)

	unitPermission, err := service.GetUnitPermissionByMemberIDAndUnitID(memberId, unit.Id)
	assert.Nil(t, err)
	assert.Equal(t, memberId, unitPermission.MemberId)
	assert.Equal(t, unit.Id, unitPermission.EntityId)
	assert.Equal(t, role, unitPermission.Role)
}

func TestUnitSerivce_AssignUnitPermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Unit{}, &models.UnitPermission{})

	memberId := "1"
	unit := models.Unit{
		Id:        "1",
		Title:     "unit title",
		CourseId:  "1",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}
	role := int32(1)
	service := NewUnitService(mockDB)
	result := service.CreateUnitNPermission(memberId, role, unit)
	assert.Nil(t, result)

	unitPermission := models.UnitPermission{
		MemberId: memberId,
		EntityId: unit.Id,
		Role:     role,
	}
	err := service.AssignUnitPermission(unitPermission)
	assert.Nil(t, err)

	unitPermission, err = service.GetUnitPermissionByMemberIDAndUnitID(memberId, unit.Id)
	assert.Nil(t, err)
	assert.Equal(t, memberId, unitPermission.MemberId)
	assert.Equal(t, unit.Id, unitPermission.EntityId)
	assert.Equal(t, role, unitPermission.Role)
}

func TestUnitSerivce_UpdateUnit(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Unit{}, &models.UnitPermission{})

	memberId := "1"
	unit := models.Unit{
		Id:        "1",
		Title:     "unit title",
		CourseId:  "1",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}
	role := int32(1)
	service := NewUnitService(mockDB)
	result := service.CreateUnitNPermission(memberId, role, unit)
	assert.Nil(t, result)

	unit.Title = "new title"
	err := service.UpdateUnit(unit)
	assert.Nil(t, err)

	units, err := service.GetUnitMemberBelongTo(memberId)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(units))
	assert.Equal(t, "new title", units[0].Title)
}

func TestUnitSerivce_DeleteUnit(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Unit{}, &models.UnitPermission{})

	memberId := "1"
	unit := models.Unit{
		Id:        "1",
		Title:     "unit title",
		CourseId:  "1",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}
	role := int32(1)
	service := NewUnitService(mockDB)
	result := service.CreateUnitNPermission(memberId, role, unit)
	assert.Nil(t, result)

	err := service.DeleteUnit(unit.Id)
	assert.Nil(t, err)

	units, err := service.GetUnitMemberBelongTo(memberId)
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(units))
}

func TestUnitService_GetUnitsByCourseID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Unit{}, &models.UnitPermission{})

	memberId := "1"
	unit := models.Unit{
		Id:        "1",
		Title:     "unit title",
		CourseId:  "1",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}
	role := int32(1)
	service := NewUnitService(mockDB)
	result := service.CreateUnitNPermission(memberId, role, unit)
	assert.Nil(t, result)

	units, err := service.GetUnitsByCourseID(unit.CourseId)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(units))
	assert.Equal(t, "unit title", units[0].Title)
}
func TestUnitService_GetUnitByMemberIDAndUnitID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Unit{}, &models.UnitPermission{})
	repo := NewUnitService(mockDB)
	memberId := "1"

	unit := models.Unit{
		Id:        "1",
		Title:     "unit title",
		Order:     1,
		Publish:   1,
		CourseId:  "1",
		CreaterId: memberId,
	}

	unit_perm := models.UnitPermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	mockDB.Create(&unit_perm)
	mockDB.Create(&unit)
	unitData, err := repo.GetUnitByMemberIDAndUnitID("1", "1")
	assert.Nil(t, err)
	assert.Equal(t, unit.Id, unitData.Id)
	assert.Equal(t, unit.Title, unitData.Title)
}
