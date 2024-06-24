package repository

import (
	"github.com/adon988/go_api_example/models"
	"gorm.io/gorm"
)

var authentication models.Authentication

type AuthRepository interface {
	DeleteAuth(id string) error
	GetAuthenticationByUsername(username string) (*gorm.DB, models.Authentication)
	CreateAuthentication(auth models.Authentication) error
}

type AuthRepositoryImpl struct {
	Db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{Db: db}
}

func (r *AuthRepositoryImpl) DeleteAuth(id string) error {
	if err := r.Db.Where("member_id = ?", id).Delete(&authentication).Error; err != nil {
		return err
	}
	return nil
}

func (r *AuthRepositoryImpl) GetAuthenticationByUsername(username string) (*gorm.DB, models.Authentication) {
	result := r.Db.First(&authentication, "username = ?", username)
	return result, authentication
}

func (r *AuthRepositoryImpl) CreateAuthentication(auth models.Authentication) error {
	return r.Db.Create(&auth).Error
}
