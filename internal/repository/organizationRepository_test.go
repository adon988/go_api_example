package repository

import (
	"fmt"
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestOrganizationRepository_CreateOrganization(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Organization{}, &models.OrganizationPermission{})

	repo := NewOrganizationRepository(mockDB)

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
	err := repo.CreateOrganization(org)
	assert.Nil(t, err)

	repoOP := NewOrganizationPermission(mockDB)
	org_perm := models.OrganizationPermission{
		MemberId: memberId,
		EntityId: org.Id,
		Role:     1,
	}
	err = repoOP.CreateOrganizationPermission(org_perm)
	assert.Nil(t, err)

	// 檢查 organization 是否被創建
	var orgs []models.Organization
	mockDB.Find(&orgs)
	assert.Equal(t, 1, len(orgs))
	assert.Equal(t, org.Id, orgs[0].Id)

	// 檢查 organization_permissions 是否被創建
	var orgPerms []models.OrganizationPermission
	mockDB.Find(&orgPerms)
	assert.Equal(t, 1, len(orgPerms))
	assert.Equal(t, memberId, orgPerms[0].MemberId)
	assert.Equal(t, org.Id, orgPerms[0].EntityId)

	fmt.Println("Result:")
	fmt.Println(orgs)
	fmt.Println(orgPerms)

	var orgM []models.Organization
	result := mockDB.Model(&models.Organization{}).Joins("JOIN organization_permissions ON organizations.id = organization_permissions.entity_id").Where("organization_permissions.member_id = ?", org_perm.MemberId).Find(&orgM)
	assert.Nil(t, result.Error)
	fmt.Println(orgM)
}

func TestOrganizationRepository_DeleteOrganization(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Organization{})

	repo := NewOrganizationRepository(mockDB)

	memberId := "1"
	Id := "1"
	org := models.Organization{
		Id:             Id,
		Title:          "org title",
		Order:          1,
		SourceLanguage: "en",
		TargetLanguage: "es",
		Publish:        1,
		CreaterId:      memberId,
	}
	err := mockDB.Create(&org)
	assert.Nil(t, err.Error)

	var orgs []models.Organization
	mockDB.Find(&orgs, "id = ?", Id)
	assert.Equal(t, 1, len(orgs))

	//刪除
	result := repo.DeleteOrganization(Id)
	assert.Nil(t, result)
	// 檢查 organization_permissions 是否被刪除
	mockDB.Find(&orgs, "id = ?", Id)
	assert.Equal(t, 0, len(orgs))
}
func TestOrganizationRepository_UpdateOrganization(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Organization{})

	repo := NewOrganizationRepository(mockDB)

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
	err := mockDB.Create(&org)
	assert.Nil(t, err.Error)

	orgUpdate := models.Organization{
		Id:    "1",
		Title: "org title update",
	}
	result := repo.UpdateOrganization(orgUpdate)
	assert.Nil(t, result)
	// 檢查 organization_permissions 是否被更新
	var orgs []models.Organization
	mockDB.Find(&orgs)
	assert.Equal(t, 1, len(orgs))
	assert.Equal(t, orgUpdate.Title, orgs[0].Title)
}

func TestOrganizationRepository_GetOrganizationByMemberID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Organization{}, &models.OrganizationPermission{})

	repo := NewOrganizationRepository(mockDB)

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

	orgs, err := repo.GetOrganizationByMemberID("1")
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)

	assert.Equal(t, orgs[0].Title, "org title")
	fmt.Println(orgs)
}

func TestOrganizationRepository_GetOrganizationByMemberIDAndOrgID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Organization{}, &models.OrganizationPermission{})
	repo := NewOrganizationRepository(mockDB)

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

	organization, err := repo.GetOrganizationByMemberIDAndOrgID("1", "1")
	assert.Nil(t, err)
	assert.Equal(t, org.Id, organization.Id)
	assert.Equal(t, org.Title, organization.Title)
}
