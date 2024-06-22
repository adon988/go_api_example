package services

import (
	"github.com/adon988/go_api_example/api/repository"
	"gorm.io/gorm"
)

type AuthService struct {
	authRepo repository.AuthRepository
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		authRepo: repository.NewAuthRepository(db),
	}
}

func (service AuthService) DeleteAuth(id string) error {
	return service.authRepo.DeleteAuth(id)
}
