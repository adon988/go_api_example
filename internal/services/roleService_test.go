package services

import (
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewRoleService(t *testing.T) {
	// Create a new mock gorm.DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	// Create a new instance of the RoleService
	service := NewRoleService(mockDB)
	// Assert that the roleRepo field is not nil
	assert.NotNil(t, service.roleRepo)
}

func TestNewRoleService_CreateRole(t *testing.T) {
	// Create a new mock gorm.DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Role{})
	// Create a new instance of the RoleService
	service := NewRoleService(mockDB)
	role := models.Role{
		Id: 1, Title: "role title",
		RoleType: "role type",
		Image:    "role.png",
	}
	err := service.CreateRole(role)
	assert.Nil(t, err)
}

func TestNewRoleService_UpdateRole(t *testing.T) {
	// Create a new mock gorm.DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Role{})
	// Create a new instance of the RoleService
	service := NewRoleService(mockDB)
	role := models.Role{
		Id: 1, Title: "role title",
		RoleType: "role type",
		Image:    "role.png",
	}
	mockDB.Create(&role)
	role.Title = "updated role title"
	err := service.UpdateRole(role)
	assert.Nil(t, err)
	assert.Equal(t, role.Title, "updated role title")
}

func TestNewRoleService_DeleteRole(t *testing.T) {
	// Create a new mock gorm.DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Role{})
	// Create a new instance of the RoleService
	service := NewRoleService(mockDB)
	role := models.Role{
		Id: 1, Title: "role title",
		RoleType: "role type",
		Image:    "role.png",
	}
	mockDB.Create(&role)
	err := service.DeleteRole(role.Id)
	assert.Nil(t, err)
}

func TestNewRoleService_GetRole(t *testing.T) {
	// Create a new mock gorm.DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Role{})
	// Create a new instance of the RoleService
	service := NewRoleService(mockDB)
	role := models.Role{
		Id: 1, Title: "role title",
		RoleType: "role type",
		Image:    "role.png",
	}
	mockDB.Create(&role)
	roles, err := service.GetRole()
	assert.Nil(t, err)
	assert.Equal(t, len(roles), 1)
	assert.Equal(t, roles[0].Title, "role title")
}
