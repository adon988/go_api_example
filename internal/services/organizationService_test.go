package services

import (
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestOrganizationService_CreateOrganizationNPermission(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Organization{}, &models.OrganizationPermission{})

	memberId := "1"
	org := models.Organization{
		Id:             "1",
		Title:          "org title",
		Order:          1,
		SourceLanguage: "en",
		TargetLanguage: "es",
		Publish:        1,
		CreaterId:      memberId,
	}
	role := "admin"
	service := NewOrganizationService(mockDB)
	result := service.CreateOrganizationNPermission(memberId, role, org)
	assert.Nil(t, result)

}

func TestOrganizationService_GetOriganization(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Organization{}, &models.OrganizationPermission{})

	memberId := "1"
	org := models.Organization{
		Id:             "1",
		Title:          "org title",
		Order:          1,
		SourceLanguage: "en",
		TargetLanguage: "es",
		Publish:        1,
		CreaterId:      memberId,
	}
	role := "admin"
	service := NewOrganizationService(mockDB)
	result := service.CreateOrganizationNPermission(memberId, role, org)
	assert.Nil(t, result)

	orgs, err := service.GetOrganization(memberId)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(orgs))
}

func TestOrganizationService_UpdateOrganization(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Organization{}, &models.OrganizationPermission{})

	memberId := "1"
	org := models.Organization{
		Id:             "1",
		Title:          "org title",
		Order:          1,
		SourceLanguage: "en",
		TargetLanguage: "es",
		Publish:        1,
		CreaterId:      memberId,
	}
	role := "admin"
	service := NewOrganizationService(mockDB)
	result := service.CreateOrganizationNPermission(memberId, role, org)
	assert.Nil(t, result)

	org.Title = "updated org title"
	err := service.UpdateOrganization(org)
	assert.Nil(t, err)
}

func TestOrganizationService_DeleteOrganization(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Organization{}, &models.OrganizationPermission{})

	memberId := "1"
	org := models.Organization{
		Id:             "1",
		Title:          "org title",
		Order:          1,
		SourceLanguage: "en",
		TargetLanguage: "es",
		Publish:        1,
		CreaterId:      memberId,
	}
	role := "admin"
	service := NewOrganizationService(mockDB)
	result := service.CreateOrganizationNPermission(memberId, role, org)
	assert.Nil(t, result)

	err := service.DeleteOrganization("1")
	assert.Nil(t, err)

}
