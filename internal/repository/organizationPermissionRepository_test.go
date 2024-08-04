package repository

import (
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestOrganizationPermissionRepository_CreateOrganizationPerimission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.OrganizationPermission{})
	repo := NewOrganizationPermission(mockDB)
	orgPerm := models.OrganizationPermission{
		Id:       "1",
		MemberId: "1",
		EntityId: "1",
		Role:     "admin",
	}
	err := repo.CreateOrganizationPermission(orgPerm)
	assert.Nil(t, err)
	// 檢查 organization_permissions 是否被創建
	var orgPerms []models.OrganizationPermission
	mockDB.Find(&orgPerms)
	assert.Equal(t, 1, len(orgPerms))
	assert.Equal(t, "1", orgPerms[0].Id)
	assert.Equal(t, "1", orgPerms[0].MemberId)
	assert.Equal(t, "1", orgPerms[0].EntityId)
}

func TestOrganizationPermissionRepository_UpdateOrganizationPermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.OrganizationPermission{})
	repo := NewOrganizationPermission(mockDB)
	orgPerm := models.OrganizationPermission{
		Id:       "1",
		MemberId: "1",
		EntityId: "1",
		Role:     "admin",
	}
	mockDB.Create(&orgPerm)
	orgPerm.Role = "member"
	err := repo.UpdateOrganizationPermission(orgPerm)
	assert.Nil(t, err)
	// 檢查 organization_permissions 是否被更新
	var orgPerms []models.OrganizationPermission
	mockDB.Find(&orgPerms)
	assert.Equal(t, 1, len(orgPerms))
	assert.Equal(t, "member", orgPerms[0].Role)
}

func TestOrganizationPermissionRepository_DeleteOrganizationPermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.OrganizationPermission{})
	repo := NewOrganizationPermission(mockDB)
	orgPerm := models.OrganizationPermission{
		Id:       "1",
		MemberId: "1",
		EntityId: "1",
		Role:     "admin",
	}
	mockDB.Create(&orgPerm)
	err := repo.DeleteOrganizationPermission("1")
	assert.Nil(t, err)
	// 檢查 organization_permissions 是否被刪除
	var orgPerms []models.OrganizationPermission
	mockDB.Find(&orgPerms)
	assert.Equal(t, 0, len(orgPerms))
}

func TestOrganizationPermissionRepository_GetOrganizationPermissionByMemberID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.OrganizationPermission{})
	repo := NewOrganizationPermission(mockDB)
	orgPerm := models.OrganizationPermission{
		Id:       "1",
		MemberId: "1",
		EntityId: "1",
		Role:     "admin",
	}
	mockDB.Create(&orgPerm)
	orgPerms, err := repo.GetOrganizationPermissionByMemberID("1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(orgPerms))
	assert.Equal(t, "1", orgPerms[0].Id)
	assert.Equal(t, "1", orgPerms[0].MemberId)
	assert.Equal(t, "1", orgPerms[0].EntityId)
}
