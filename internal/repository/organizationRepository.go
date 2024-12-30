package repository

import (
	"fmt"

	"github.com/adon988/go_api_example/internal/models"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	CreateOrganization(organization models.Organization) error
	UpdateOrganization(organization models.Organization) error
	DeleteOrganization(id string) error
	GetOrganizationByMemberID(member_id string) ([]models.Organization, error)
	GetOrganizationByMemberIDAndOrgID(member_id string, organization_id string) (models.Organization, error)
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &OrganizationRepositoryImpl{DB: db}
}

type OrganizationRepositoryImpl struct {
	DB *gorm.DB
}

// Implement the interface methods
func (r *OrganizationRepositoryImpl) CreateOrganization(organization models.Organization) error {
	err := r.DB.Create(&organization)

	if err != nil {
		return err.Error
	}

	return nil
}

func (r *OrganizationRepositoryImpl) UpdateOrganization(organization models.Organization) error {
	err := r.DB.Where("id = ?", organization.Id).Updates(organization)
	if err.Error != nil {
		return fmt.Errorf("oh! %s", err.Error)
	}
	if err.RowsAffected == 0 {
		return fmt.Errorf("no organization found with id: %s", organization.Id)
	}
	return nil
}

func (r *OrganizationRepositoryImpl) DeleteOrganization(id string) error {
	result := r.DB.Delete(&models.Organization{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no organization found with id: %s", id)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no organization found with id: %s", id)
	}
	return nil
}

func (r *OrganizationRepositoryImpl) GetOrganizationByMemberID(member_id string) ([]models.Organization, error) {

	var orgs []models.Organization
	err := r.DB.Model(&models.Organization{}).Joins("JOIN organization_permissions ON organizations.id = organization_permissions.entity_id").Where("organization_permissions.member_id = ?", member_id).Find(&orgs)
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected == 0 {
		return nil, fmt.Errorf("no organization found with member id: %s", member_id)
	}
	return orgs, nil
}

func (r *OrganizationRepositoryImpl) GetOrganizationByMemberIDAndOrgID(member_id string, organization_id string) (models.Organization, error) {
	var organization models.Organization

	//Member should with eitity permission or entity is publis
	err := r.DB.Model(&models.Organization{}).Joins("JOIN organization_permissions ON organizations.id = organization_permissions.entity_id").Where("(organization_permissions.member_id = ? OR organizations.publish = 1) AND organizations.id = ?", member_id, organization_id).Find(&organization)
	if err.Error != nil {
		return models.Organization{}, err.Error
	}
	if err.RowsAffected == 0 {
		return models.Organization{}, fmt.Errorf("no organization found with member id: %s and organization id: %s", member_id, organization_id)
	}
	return organization, nil
}
