package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/adon988/go_api_example/api/repository"
	"github.com/adon988/go_api_example/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewMemberService(t *testing.T) {
	// Create a new mock gorm.DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	// Create a new instance of the MemberService
	service := NewMemberService(mockDB)
	// Assert that the memberRepo field is not nil
	assert.NotNil(t, service.memberRepo)
}

type MemberRepositoryMock struct {
	DB *gorm.DB
}

func MockMemberRepository(db *gorm.DB) repository.MemberRepository {
	return &MemberRepositoryMock{DB: db}
}

func (r *MemberRepositoryMock) GetMemberInfo(id string) (*models.Member, error) {
	// Mock the behavior of the memberRepo.GetMemberInfo method
	name := "John Doe"
	birthday := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	email := "asdf@adsf.asdf"
	if id == "1" {
		return &models.Member{Id: "1", Name: &name, Birthday: &birthday, Email: &email}, nil
	}
	return nil, fmt.Errorf("member not found")
}

func TestMemberService_GetMemberInfo(t *testing.T) {
	// Create a new mock gorm.DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	// Mock the behavior of the DB.First method
	name := "John Doe"
	birthday := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	email := "john@example.com"
	gender := 1
	mockMember := &models.Member{Id: "1", Name: &name, Birthday: &birthday, Email: &email, Gender: &gender}

	// migrate schema
	mockDB.AutoMigrate(&models.Member{})
	// auto insert data to db with mock member
	mockDB.Create(&mockMember)

	// Create a new instance of the MemberService
	service := NewMemberService(mockDB)

	// Call the GetMemberInfo method with a valid ID
	member, err := service.GetMemberInfo("1")

	// Assert that the returned member is the expected member
	assert.Equal(t, mockMember.Id, member.Id)
	assert.Equal(t, mockMember.Name, member.Name)
	assert.Equal(t, mockMember.Birthday, member.Birthday)
	assert.Equal(t, mockMember.Email, member.Email)
	assert.Equal(t, mockMember.Gender, member.Gender)
	// Assert that the error is nil
	assert.Nil(t, err)

	// Call the GetMemberInfo method with an invalid ID
	member, err = service.GetMemberInfo("2")

	// Assert that the returned member is nil
	assert.Nil(t, member)
	// Assert that the error is not nil
	assert.NotNil(t, err)
}
