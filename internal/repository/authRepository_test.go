package repository

import (
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestAuthRepositoryImpl_CreateAuthentication(t *testing.T) {
	// Create a new mock DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	// migrare schema
	mockDB.AutoMigrate(&models.Authentication{})
	// Create a new instance of the MemberRepositoryImpl
	repo := NewAuthRepository(mockDB)
	username := "john"
	password, _ := bcrypt.GenerateFromPassword([]byte("mypassword"), bcrypt.DefaultCost)
	memberId, _ := utils.GenId()

	Type := "ApikeyAuth"
	oriAuth := &models.Authentication{Username: username, Password: password, MemberId: memberId, Type: &Type}
	err := repo.CreateAuthentication(*oriAuth)
	assert.NoError(t, err)

	var auth models.Authentication
	mockDB.First(&auth, "username = ?", username)
	assert.Equal(t, auth.Username, username)
}

func TestAuthRepositoryImpl_GetAuthenticationByUsername(t *testing.T) {
	// Create a new mock DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	// migrare schema
	mockDB.AutoMigrate(&models.Authentication{})
	// Create a new instance of the MemberRepositoryImpl
	repo := NewAuthRepository(mockDB)
	username := "john"
	password, _ := bcrypt.GenerateFromPassword([]byte("mypassword"), bcrypt.DefaultCost)
	memberId, _ := utils.GenId()

	Type := "ApikeyAuth"
	oriAuth := &models.Authentication{Username: username, Password: password, MemberId: memberId, Type: &Type}
	mockDB.Create(&oriAuth)
	authResult, result := repo.GetAuthenticationByUsername(username)
	assert.Equal(t, authResult.RowsAffected, int64(1))
	assert.Equal(t, result.Username, username)
}

func TestAuthRepositoryImpl_DeleteAuth(t *testing.T) {
	// Create a new mock DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	// migrare schema
	mockDB.AutoMigrate(&models.Authentication{})
	// Create a new instance of the MemberRepositoryImpl
	repo := NewAuthRepository(mockDB)

	username := "john"
	password, _ := bcrypt.GenerateFromPassword([]byte("mypassword"), bcrypt.DefaultCost)
	memberId, _ := utils.GenId()

	Type := "ApikeyAuth"
	oriAuth := &models.Authentication{Username: username, Password: password, MemberId: memberId, Type: &Type}
	mockDB.Create(&oriAuth)

	err := repo.DeleteAuth(memberId)
	assert.Nil(t, err)

}
