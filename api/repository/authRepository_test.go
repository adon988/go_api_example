package repository

import (
	"testing"

	"github.com/adon988/go_api_example/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestAuthRepositoryImpl_DeleteAuth(t *testing.T) {
	// Create a new mock DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	// migrare schema
	mockDB.AutoMigrate(&models.Authentication{})
	// Create a new instance of the MemberRepositoryImpl
	repo := NewAuthRepository(mockDB)

	username := "john"
	password, _ := bcrypt.GenerateFromPassword([]byte("mypassword"), bcrypt.DefaultCost)
	memberId := "1"
	Type := "ApikeyAuth"
	oriAuth := &models.Authentication{Username: username, Password: password, MemberId: memberId, Type: &Type}
	mockDB.Create(&oriAuth)

	err := repo.DeleteAuth(memberId)
	assert.Nil(t, err)

}
