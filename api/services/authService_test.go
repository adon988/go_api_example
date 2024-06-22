package services

import (
	"testing"

	"github.com/adon988/go_api_example/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewService(t *testing.T) {
	// Create a new mock gorm.DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	// Create a new instance of the MemberService
	service := NewAuthService(mockDB)
	// Assert that the memberRepo field is not nil
	assert.NotNil(t, service.authRepo)
}

func TestAuthService_DeleteAuth(t *testing.T) {
	// Create a new mock gorm.DB instance
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	// migrate schema
	mockDB.AutoMigrate(&models.Authentication{})
	// Create default member data
	username := "john"
	password, _ := bcrypt.GenerateFromPassword([]byte("mypassword"), bcrypt.DefaultCost)
	memberId := "1"
	Type := "ApikeyAuth"
	oriAuth := &models.Authentication{Username: username, Password: password, MemberId: memberId, Type: &Type}
	mockDB.Create(&oriAuth)

	// Mock the behavior of the DB.First method
	service := NewAuthService(mockDB)
	err := service.DeleteAuth(memberId)
	assert.Nil(t, err)
}
