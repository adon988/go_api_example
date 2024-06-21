package services

import (
	"github.com/adon988/go_api_example/api/repository"
	"github.com/adon988/go_api_example/models"
	"gorm.io/gorm"
)

type MemberService struct {
	memberRepo repository.MemberRepository
}

func NewMemberService(db *gorm.DB) *MemberService {
	return &MemberService{
		memberRepo: repository.NewMemberRepository(db),
	}
}

func (service MemberService) GetMemberInfo(id string) (*models.Member, error) {
	return service.memberRepo.GetMemberInfo(id)
}
