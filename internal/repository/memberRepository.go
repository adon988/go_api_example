package repository

import (
	"github.com/adon988/go_api_example/internal/models"
	"gorm.io/gorm"
)

type MemberRepository interface {
	GetMemberInfo(id string) (*models.Member, error)
	GetMembersWithRoles(id string) (*models.Member, error)
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
func (r *MemberRepositoryImpl) GetMemberInfo(id string) (*models.Member, error) {
	if err := r.DB.First(&member, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *MemberRepositoryImpl) GetMembersWithRoles(id string) (*models.Member, error) {
	var member models.Member
	err := r.DB.Joins("Role").First(&member, "members.id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &member, nil
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
