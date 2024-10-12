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

func TestOrganizationService_GetorganizationByMemberIDAndOrganizationID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Organization{}, &models.OrganizationPermission{})

	org := models.Organization{
		Id:             "1",
		Title:          "org title",
		Order:          1,
		SourceLanguage: "en",
		TargetLanguage: "es",
		Publish:        1,
		CreaterId:      "1",
	}
	org_perm := models.OrganizationPermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	mockDB.Create(&org)
	mockDB.Create(&org_perm)

	service := NewOrganizationService(mockDB)
	organization, err := service.GetOrganizationByMemberIDAndOrganizationID("1", "1")
	assert.Nil(t, err)
	assert.Equal(t, org.Id, organization.Id)
	assert.Equal(t, org.Title, organization.Title)
}
