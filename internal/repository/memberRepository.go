package repository

import (
	"fmt"

	"github.com/adon988/go_api_example/internal/models"
	"gorm.io/gorm"
)

type MemberRepository interface {
	GetMemberInfo(id string) (models.Member, error)
	GetMembersWithRoles(id string) (models.Member, error)
	GetMembersValid(member_id []string) ([]string, error)
	CreateMember(id string) error
	UpdateMember(id string, data models.Member) error
	DeleteMember(id string) error
}

type MemberRepositoryImpl struct {
	DB *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &MemberRepositoryImpl{DB: db}
}

var member models.Member

// GetMemberInfo implements MemberRepository.
func (r *MemberRepositoryImpl) GetMemberInfo(id string) (models.Member, error) {
	var member models.Member
	if err := r.DB.Where("id = ?", id).First(&member).Error; err != nil {
		return member, err
	}
	return member, nil
}

func (r *MemberRepositoryImpl) GetMembersWithRoles(id string) (models.Member, error) {
	var member models.Member
	err := r.DB.Joins("Role").First(&member, "members.id = ?", id).Error
	if err != nil {
		return member, err
	}

	return member, nil
}

func (r *MemberRepositoryImpl) GetMembersValid(member_id []string) ([]string, error) {
	var validMemberIDs []string
	if err := r.DB.Model(&models.Member{}).Where("id IN (?)", member_id).Pluck("id", &validMemberIDs).Error; err != nil {
		return []string{}, err
	}

	if len(validMemberIDs) == 0 {
		return []string{}, fmt.Errorf("no valid members found")
	}

	return validMemberIDs, nil
}

func (r *MemberRepositoryImpl) CreateMember(id string) error {
	member := &models.Member{
		Id: id,
	}
	err := r.DB.Create(&member)
	return err.Error
}

func (r *MemberRepositoryImpl) UpdateMember(id string, data models.Member) error {
	if err := r.DB.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

// DeleteMember implements MemberRepository.
func (r *MemberRepositoryImpl) DeleteMember(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&member).Error; err != nil {
		return err
	}
	return nil
}
