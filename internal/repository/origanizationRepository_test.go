package repository

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestTest(t *testing.T) {
	m := []models.Permission{
		{
			UserID: "1",
			Role:   "admin",
		},
		{
			UserID: "2",
			Role:   "admin",
		},
	}
	m = append(m, models.Permission{UserID: "3", Role: "writer"})
	m = append(m, models.Permission{UserID: "3", Role: "writer"})
	j := make(map[string]string)
	var o []models.Permission
	for _, v := range m {
		if _, ok := j[v.UserID+v.Role]; ok {
			continue
		}
		j[v.UserID+v.Role] = "ok"
		o = append(o, v)
	}
	fmt.Println(m)
	fmt.Println(o)

	m_struc_to_json, _ := json.Marshal(m)
	perm := string(m_struc_to_json)
	fmt.Println(perm)

	str := `[{"user_id":"1","role":"admin"},{"user_id":"3","role":"writer"}]`
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(m)

	assert.Equal(t, 1, 1)
}

func TestOriganizationRepository_CreateOriganization(t *testing.T) {

	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	mockDB.AutoMigrate(&models.Organization{})
	repo := NewOriganizationRepository(mockDB, false)
	permission_json, _ := json.Marshal(
		[]models.Permission{
			{
				UserID: "1",
				Role:   "admin",
			},
			{
				UserID: "2",
				Role:   "admin",
			},
		},
	)
	permission_json_str := string(permission_json)

	data := models.Organization{
		Id:              "1",
		Title:           "origanization title",
		Order:           1,
		SourceLangeuage: "en",
		TargetLanguage:  "vi",
		CreaterId:       "1",
		Permissions:     permission_json_str,
	}
	err := repo.CreateOriganization(data)
	assert.Nil(t, err)
}

func TestOriganizationRepository_DeleteOriganization(t *testing.T) {
	assert.Equal(t, 1, 1)
}
func TestOriganizationRepository_UpdateOriganization(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestOriganizationRepository_GetOriganization(t *testing.T) {

	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	mockDB.AutoMigrate(&models.Organization{})
	repo := NewOriganizationRepository(mockDB, false)
	permission_json, _ := json.Marshal(
		[]models.Permission{
			{
				UserID: "1",
				Role:   "admin",
			},
			{
				UserID: "2",
				Role:   "admin",
			},
		},
	)
	permission_json_str := string(permission_json)

	data := models.Organization{
		Id:              "1",
		Title:           "origanization title",
		Order:           1,
		SourceLangeuage: "en",
		TargetLanguage:  "vi",
		CreaterId:       "1",
		Permissions:     permission_json_str,
	}
	err := repo.CreateOriganization(data)
	assert.Nil(t, err)

	res, _ := repo.GetOriganizationByMemberID(data.Id)
	assert.Equal(t, data.Permissions, res[0].Permissions)

}
