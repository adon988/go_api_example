// repository/memberRepository_test.go
package repository

import (
	"testing"
	"time"

	"github.com/adon988/go_api_example/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMemberRepositoryImpl_GetMemberInfo(t *testing.T) {
	// Create a new mock DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	// Create a new instance of the MemberRepositoryImpl
	repo := NewMemberRepository(mockDB)

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

	// Call the GetMemberInfo method with a valid ID
	member, err := repo.GetMemberInfo("1")

	// Assert that the returned member is the expected member
	assert.Equal(t, mockMember.Id, member.Id)
	assert.Equal(t, mockMember.Name, member.Name)
	assert.Equal(t, mockMember.Birthday, member.Birthday)
	assert.Equal(t, mockMember.Email, member.Email)
	assert.Equal(t, mockMember.Gender, member.Gender)

	// Assert that the error is nil
	assert.Nil(t, err)

	// Call the GetMemberInfo method with an invalid ID
	member, err = repo.GetMemberInfo("2")

	// Assert that the returned member is nil
	assert.Nil(t, member)

	// Assert that the error is not nil
	assert.NotNil(t, err)
}
