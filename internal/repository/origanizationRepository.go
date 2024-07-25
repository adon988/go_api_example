package repository

import (
	"github.com/adon988/go_api_example/internal/models"
	"gorm.io/gorm"
)

type OriganizationRepository interface {
	CreateOriganization(origanization models.Organization) error
	UpdateOriganization(origanization models.Organization) error
	DeleteOriganization(id string) error
	GetOriganizationByMemberID(id string) ([]models.Organization, error)
}

type OriganizationRepositoryImpl struct {
	DB *gorm.DB
}

func NewOriganizationRepository(db *gorm.DB, support_json bool) OriganizationRepository {
	return &OriganizationRepositoryImpl{DB: db}
}

// Implement the interface methods
func (r *OriganizationRepositoryImpl) CreateOriganization(origanization models.Organization) error {
	return nil
}
func (r *OriganizationRepositoryImpl) UpdateOriganization(origanization models.Organization) error {
	return nil
}

func (r *OriganizationRepositoryImpl) DeleteOriganization(id string) error {
	return nil
}

func (r *OriganizationRepositoryImpl) GetOriganizationByMemberID(id string) ([]models.Organization, error) {
	return nil, nil
}
