// repository/memberRepository_test.go
package repository

import (
	"testing"
	"time"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMemberRepositoryImpl_CreateMember(t *testing.T) {
	// Create a new mock DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	// migrare schema
	mockDB.AutoMigrate(&models.Member{})
	// Create a new instance of the MemberRepositoryImpl
	repo := NewMemberRepository(mockDB)

	id := "1"
	err := repo.CreateMember(id)
	assert.Nil(t, err)

}

func TestMemberRepositoryImpl_DeleteMember(t *testing.T) {
	// Create a new mock DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	// migrare schema
	mockDB.AutoMigrate(&models.Member{})
	// Create a new instance of the MemberRepositoryImpl
	repo := NewMemberRepository(mockDB)
	id := "1"
	name := "John Doe"
	birthday := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	email := "john@example.com"
	gender := int32(1)
	origMember := &models.Member{Id: id, Name: &name, Birthday: &birthday, Email: &email, Gender: &gender}
	mockDB.Create(&origMember)

	err := repo.DeleteMember(id)
	assert.Nil(t, err)

}

func TestMemberRepositoryImpl_UPdateMember(t *testing.T) {
	// Create a new mock DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	// migrate schema
	mockDB.AutoMigrate(&models.Member{})

	// Create a new instance of the MemberRepositoryImpl
	repo := NewMemberRepository(mockDB)
	id := "1"
	name := "John Doe"
	birthday := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	email := "john@example.com"
	gender := int32(1)
	origMember := &models.Member{Id: id, Name: &name, Birthday: &birthday, Email: &email, Gender: &gender}
	mockDB.Create(&origMember)

	name = "Jane Doe2"
	birthday = time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	email = "john2@example.com"
	gender = int32(0)
	mockMember := &models.Member{Name: &name, Birthday: &birthday, Email: &email, Gender: &gender}
	err := repo.UpdateMember(id, *mockMember)
	assert.Nil(t, err)

	// Call the GetMemberInfo method with a valid ID
	member, err := repo.GetMemberInfo(id)
	assert.Nil(t, err)
	assert.Equal(t, mockMember.Name, member.Name)
	assert.Equal(t, mockMember.Birthday, member.Birthday)
	assert.Equal(t, mockMember.Email, member.Email)
	assert.Equal(t, mockMember.Gender, member.Gender)

}

func TestMemberRepositoryImpl_GetMemberValid(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Member{})
	repo := NewMemberRepository(mockDB)

	id := "1"
	name := "John Doe"
	birthday := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	email := "john@example.com"
	gender := int32(1)
	mockMember := &models.Member{Id: id, Name: &name, Birthday: &birthday, Email: &email, Gender: &gender}
	mockDB.Create(&mockMember)

	members, err := repo.GetMembersValid([]string{"1"})

	assert.Nil(t, err)
	assert.Equal(t, members, []string{"1"})

	id = "2"
	name = "John Doe2"
	birthday = time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	email = "john2@example.com"
	gender = int32(2)
	mockMember2 := &models.Member{Id: id, Name: &name, Birthday: &birthday, Email: &email, Gender: &gender}
	mockDB.Create(&mockMember2)

	members, err = repo.GetMembersValid([]string{"1", "2"})
	assert.Nil(t, err)
	assert.Equal(t, members, []string{"1", "2"})

	members, err = repo.GetMembersValid([]string{"1", "3"})
	assert.Nil(t, err)
	assert.Equal(t, members, []string{"1"})

	members, err = repo.GetMembersValid([]string{"4", "3"})
	assert.NotNil(t, err)
	assert.Equal(t, members, []string{})
}

func TestMemberRepositoryImpl_GetMemberInfo(t *testing.T) {
	// Create a new mock DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	// Create a new instance of the MemberRepositoryImpl
	repo := NewMemberRepository(mockDB)

	id := "1"
	name := "John Doe"
	birthday := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	email := "john@example.com"
	gender := int32(1)
	mockMember := &models.Member{Id: id, Name: &name, Birthday: &birthday, Email: &email, Gender: &gender}

	// migrate schema
	mockDB.AutoMigrate(&models.Member{})
	// auto insert data to db with mock member
	mockDB.Create(&mockMember)

	// Call the GetMemberInfo method with a valid ID
	member, err := repo.GetMemberInfo(id)

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

	// Assert that the error is not nil
	assert.NotNil(t, err)
}
