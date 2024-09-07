package repository

import (
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUnitPermissionRepository_CreateUnitPermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.UnitPermission{})
	repo := NewUnitPermissionRepository(mockDB)
	unitPerm := models.UnitPermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	err := repo.CreateUnitPermission(unitPerm)
	assert.Nil(t, err)
	// 檢查 unit_permissions 是否被建立
	var unitPerms []models.UnitPermission
	mockDB.Find(&unitPerms)

	assert.Equal(t, 1, len(unitPerms))
	assert.Equal(t, "1", unitPerms[0].MemberId)
	assert.Equal(t, "1", unitPerms[0].EntityId)
	// test role not found
	unitPerm = models.UnitPermission{
		MemberId: "1",
		EntityId: "1",
		Role:     4,
	}
	err = repo.CreateUnitPermission(unitPerm)
	assert.Equal(t, "role not found", err.Error())
}

func TestUnitPermissionRepository_GetUnitByMemberIDAndUnitID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.UnitPermission{})
	repo := NewUnitPermissionRepository(mockDB)
	unitPerm := models.UnitPermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	//creatre
	err := repo.AssignUnitPermission(unitPerm)
	assert.Nil(t, err)
	unitPerms, err := repo.GetUnitPermissionByMemberIDAndUnitID("1", "1")
	assert.Nil(t, err)
	assert.Equal(t, "1", unitPerms.MemberId)
	assert.Equal(t, "1", unitPerms.EntityId)
	assert.Equal(t, int32(1), unitPerm.Role)
	//update
	unitPerm2 := models.UnitPermission{
		MemberId: "1",
		EntityId: "1",
		Role:     2,
	}
	err = repo.AssignUnitPermission(unitPerm2)
	assert.Nil(t, err)
	unitPerms, err = repo.GetUnitPermissionByMemberIDAndUnitID("1", "1")
	assert.Nil(t, err)
	assert.Equal(t, int32(2), unitPerms.Role)
}

func TestUnitPermissionRepository_UpdateUnitPermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.UnitPermission{})
	repo := NewUnitPermissionRepository(mockDB)
	unitPerm := models.UnitPermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	mockDB.Create(&unitPerm)
	unitPerm.Role = 2
	err := repo.UpdateUnitPermission(unitPerm)
	assert.Nil(t, err)
	// 檢查 unit_permissions 是否被更新
	var unitPerms []models.UnitPermission
	mockDB.Find(&unitPerms)
	assert.Equal(t, 1, len(unitPerms))
	assert.Equal(t, int32(2), unitPerms[0].Role)
}

func TestUnitPermissionRepository_DeleteUnitPermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.UnitPermission{})
	repo := NewUnitPermissionRepository(mockDB)
	unitPerm := models.UnitPermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	mockDB.Create(&unitPerm)
	err := repo.DeleteUnitPermission("1")
	assert.Nil(t, err)
	// 檢查 unit_permissions 是否被刪除
	var unitPerms []models.UnitPermission
	mockDB.Find(&unitPerms)
	assert.Equal(t, 0, len(unitPerms))
}

func TestUnitPermissionRepository_AssignUnitPermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.UnitPermission{})
	repo := NewUnitPermissionRepository(mockDB)
	unitPerm := models.UnitPermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	//create
	err := repo.AssignUnitPermission(unitPerm)
	assert.Nil(t, err)
	// 檢查 unit_permissions 是否被建立
	var unitPerms []models.UnitPermission
	mockDB.Find(&unitPerms)
	assert.Equal(t, 1, len(unitPerms))
	assert.Equal(t, "1", unitPerms[0].MemberId)
	assert.Equal(t, "1", unitPerms[0].EntityId)
	// test role not found
	unitPerm = models.UnitPermission{
		MemberId: "1",
		EntityId: "1",
		Role:     4,
	}
	err = repo.AssignUnitPermission(unitPerm)
	assert.Equal(t, "role not found", err.Error())
	//update
	unitPerm2 := models.UnitPermission{
		MemberId: "1",
		EntityId: "1",
		Role:     2,
	}
	err = repo.AssignUnitPermission(unitPerm2)
	assert.Nil(t, err)
	// 檢查 unit_permissions 是否被更新
	mockDB.Find(&unitPerms)
	assert.Equal(t, 1, len(unitPerms))
	assert.Equal(t, int32(2), unitPerms[0].Role)
}
