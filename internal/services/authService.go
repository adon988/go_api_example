package services

import (
	"errors"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/repository"
	"gorm.io/gorm"
)

type AuthService struct {
	authRepo   repository.AuthRepository
	memberRepo repository.MemberRepository
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		authRepo:   repository.NewAuthRepository(db),
		memberRepo: repository.NewMemberRepository(db),
	}
}

func (service AuthService) DeleteAuth(id string) error {
	return service.authRepo.DeleteAuth(id)
}

func (service AuthService) Register(auth models.Authentication) error {

	result, _ := service.authRepo.GetAuthenticationByUsername(auth.Username)
	if result.RowsAffected > 0 {
		return errors.New("account already exists(:0)")
	}

	err := service.authRepo.CreateAuthentication(auth)
	if err != nil {
		return errors.New("failed to create account(:1)")
	}

	err = service.memberRepo.CreateMember(auth.MemberId)
	if err != nil {
		return errors.New("failed to create account(:2)")
	}

	return nil
}
