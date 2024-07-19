package repository

import (
	"testing"

	"github.com/adon988/go_api_example/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestRoleRepositoryImpl_CreateRole(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Role{})
	repo := NewRoleRepository(mockDB)
	role := models.Role{
		Id: 1, Title: "role title",
		RoleType: "role type",
		Image:    "role.png",
	}
	err := repo.CreateRole(role)
	assert.Nil(t, err)

}

func TestRoleRepositoryImpl_UpdateRole(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Role{})
	repo := NewRoleRepository(mockDB)
	role := models.Role{
		Id: 1, Title: "role title",
		RoleType: "role type",
		Image:    "role.png",
	}
	mockDB.Create(&role)
	role.Title = "updated role title"
	err := repo.UpdateRole(role)
	assert.Nil(t, err)
	assert.Equal(t, role.Title, "updated role title")
}

func TestRoleRepositoryImpl_GetRole(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Role{})
	repo := NewRoleRepository(mockDB)
	role := models.Role{
		Id: 1, Title: "role title",
		RoleType: "role type",
		Image:    "role.png",
	}
	mockDB.Create(&role)
	roles, err := repo.GetRole()
	assert.Nil(t, err)
	assert.Equal(t, len(roles), 1)
	assert.Equal(t, roles[0].Title, "role title")
}

func TestRoleRepositoryImpl_DeleteRo(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Role{})
	repo := NewRoleRepository(mockDB)
	role := models.Role{
		Id: 1, Title: "role title",
		RoleType: "role type",
		Image:    "role.png",
	}
	mockDB.Create(&role)
	err := repo.DeleteRole(1)
	assert.Nil(t, err)
}
