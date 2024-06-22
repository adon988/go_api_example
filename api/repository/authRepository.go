package repository

import (
	"github.com/adon988/go_api_example/models"
	"gorm.io/gorm"
)

var authentication models.Authentication

type AuthRepository interface {
	DeleteAuth(id string) error
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
