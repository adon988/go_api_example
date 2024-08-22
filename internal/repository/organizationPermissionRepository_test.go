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
		Role:     1,
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

	orgPerm = models.OrganizationPermission{
		Id:       "2",
		MemberId: "1",
		EntityId: "1",
		Role:     4,
	}
	err = repo.CreateOrganizationPermission(orgPerm)
	assert.Equal(t, "role not found", err.Error())
}

func TestOrganizationPermissionRepository_UpdateOrganizationPermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.OrganizationPermission{})
	repo := NewOrganizationPermission(mockDB)
	orgPerm := models.OrganizationPermission{
		Id:       "1",
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	mockDB.Create(&orgPerm)
	orgPerm.Role = 2
	err := repo.UpdateOrganizationPermission(orgPerm)
	assert.Nil(t, err)
	// 檢查 organization_permissions 是否被更新
	var orgPerms []models.OrganizationPermission
	mockDB.Find(&orgPerms)
	assert.Equal(t, 1, len(orgPerms))
	assert.Equal(t, int32(2), orgPerms[0].Role)
}

func TestOrganizationPermissionRepository_DeleteOrganizationPermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.OrganizationPermission{})
	repo := NewOrganizationPermission(mockDB)
	orgPerm := models.OrganizationPermission{
		Id:       "1",
		MemberId: "1",
		EntityId: "1",
		Role:     1,
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
		Role:     1,
	}
	mockDB.Create(&orgPerm)
	orgPerms, err := repo.GetOrganizationPermissionByMemberID("1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(orgPerms))
	assert.Equal(t, "1", orgPerms[0].Id)
	assert.Equal(t, "1", orgPerms[0].MemberId)
	assert.Equal(t, "1", orgPerms[0].EntityId)
}

func TestOrganizationPermissionRepository_GetOrganizationPermissionByOrganizationIDAndMemberID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.OrganizationPermission{})
	repo := NewOrganizationPermission(mockDB)
	orgPerm := models.OrganizationPermission{
		Id:       "1",
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	mockDB.Create(&orgPerm)
	orgPerms, err := repo.GetOrganizationPermissionByOrganizationIDAndMemberID("1", "1")
	assert.Nil(t, err)
	assert.Equal(t, "1", orgPerms.Id)
	assert.Equal(t, "1", orgPerms.MemberId)
	assert.Equal(t, "1", orgPerms.EntityId)

	orgPerms, err = repo.GetOrganizationPermissionByOrganizationIDAndMemberID("1", "2")
	assert.NotNil(t, err)
}
