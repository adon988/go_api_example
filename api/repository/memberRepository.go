package repository

import (
	"github.com/adon988/go_api_example/models"
	"gorm.io/gorm"
)

type MemberRepository interface {
	GetMemberInfo(id string) (*models.Member, error)
}

type MemberRepositoryImpl struct {
	DB *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &MemberRepositoryImpl{DB: db}
}

// GetMemberInfo implements MemberRepository.
func (r *MemberRepositoryImpl) GetMemberInfo(id string) (*models.Member, error) {
	var member models.Member
	if err := r.DB.First(&member, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &member, nil
}
