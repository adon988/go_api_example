package services

import (
	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/repository"
	"gorm.io/gorm"
)

type MemberService struct {
	memberRepo repository.MemberRepository
	authRepo   repository.AuthRepository
}

func NewMemberService(db *gorm.DB) *MemberService {
	return &MemberService{
		memberRepo: repository.NewMemberRepository(db),
		authRepo:   repository.NewAuthRepository(db),
	}
}

func (service MemberService) GetMemberInfo(id string) (models.Member, error) {
	return service.memberRepo.GetMemberInfo(id)
}

func (service MemberService) GetMembersWithRoles(id string) (models.Member, error) {
	return service.memberRepo.GetMembersWithRoles(id)
}

func (service MemberService) GetValidMembers(member_id []string) ([]string, error) {
	members, err := service.memberRepo.GetMembersValid(member_id)
	if err != nil {
		return []string{}, err
	}

	return members, nil
}

func (service MemberService) UpdateMember(id string, data models.Member) error {
	return service.memberRepo.UpdateMember(id, data)
}

func (service MemberService) DeleteMemberAndAuth(id string) error {
	if err := service.authRepo.DeleteAuth(id); err != nil {
		return err
	}
	if err := service.memberRepo.DeleteMember(id); err != nil {
		return err
	}
	return nil
}
