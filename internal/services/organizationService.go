package services

import (
	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/repository"
	"gorm.io/gorm"
)

type OrganizationServiceInterface interface {
	CreateOrganizationNPermission(member_id string, role string, organization models.Organization) error
	GetOrganization(member_id string) ([]models.Organization, error)
	UpdateOrganization(organization models.Organization) error
	DeleteOrganization(id string) error
}

func NewOrganizationService(db *gorm.DB) OrganizationServiceInterface {
	return &OrganizationService{
		organization:           repository.NewOrganizationRepository(db),
		organizationPermission: repository.NewOrganizationPermission(db),
	}
}

type OrganizationService struct {
	organization           repository.OrganizationRepository
	organizationPermission repository.OrganizationPermissionRepository
}

func (service OrganizationService) CreateOrganizationNPermission(member_id string, role string, organization models.Organization) error {

	organizationPermission := models.OrganizationPermission{
		MemberId: member_id,
		EntityId: organization.Id,
		Role:     role,
	}

	result := service.organizationPermission.CreateOrganizationPermission(organizationPermission)
	if result != nil {
		return result
	}

	err := service.organization.CreateOrganization(member_id, organization)
	if err != nil {
		return err
	}
	return nil
}

func (service OrganizationService) GetOrganization(member_id string) ([]models.Organization, error) {
	return service.organization.GetOrganizationByMemberID(member_id)
}

func (service OrganizationService) UpdateOrganization(organization models.Organization) error {
	//這裡要再檢查 role 是否有 admin 權限
	return service.organization.UpdateOrganization(organization)
}

func (service OrganizationService) DeleteOrganization(id string) error {
	//這裡要再檢查 role 是否有 admin 權限
	if err := service.organizationPermission.DeleteOrganizationPermission(id); err != nil {
		return err
	}
	if err := service.organization.DeleteOrganization(id); err != nil {
		return err
	}
	return nil
}
