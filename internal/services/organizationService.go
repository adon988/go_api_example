package services

import (
	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/repository"
	"github.com/adon988/go_api_example/internal/utils"
	"gorm.io/gorm"
)

type OrganizationServiceInterface interface {
	CreateOrganizationNPermission(member_id string, role string, organization models.Organization) error
	GetOrganization(member_id string) ([]models.Organization, error)
	GetOrganizationByMemberIDAndOrganizationID(member_id string, organization_id string) (models.Organization, error)
	UpdateOrganization(organization models.Organization) error
	DeleteOrganization(id string) error
	GetOrganizationPermissionByOrganizationIDAndMemberID(member_id string, organization_id string) (models.OrganizationPermission, error)
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
	orgPerId, _ := utils.GenId()
	organizationPermission := models.OrganizationPermission{
		Id:       orgPerId,
		MemberId: member_id,
		EntityId: organization.Id,
		Role:     role,
	}

	result := service.organizationPermission.CreateOrganizationPermission(organizationPermission)
	if result != nil {
		return result
	}

	err := service.organization.CreateOrganization(organization)
	if err != nil {
		return err
	}
	return nil
}

func (service OrganizationService) GetOrganizationByMemberIDAndOrganizationID(member_id string, organization_id string) (models.Organization, error) {
	return service.organization.GetOrganizationByMemberIDAndOrgID(member_id, organization_id)
}
func (service OrganizationService) GetOrganization(member_id string) ([]models.Organization, error) {
	return service.organization.GetOrganizationByMemberID(member_id)
}

func (service OrganizationService) UpdateOrganization(organization models.Organization) error {
	return service.organization.UpdateOrganization(organization)
}

func (service OrganizationService) DeleteOrganization(id string) error {
	if err := service.organizationPermission.DeleteOrganizationPermission(id); err != nil {
		return err
	}
	if err := service.organization.DeleteOrganization(id); err != nil {
		return err
	}
	return nil
}

func (service OrganizationService) GetOrganizationPermissionByOrganizationIDAndMemberID(member_id string, organization_id string) (models.OrganizationPermission, error) {
	result, err := service.organizationPermission.GetOrganizationPermissionByOrganizationIDAndMemberID(member_id, organization_id)
	if err != nil {
		return models.OrganizationPermission{}, err
	}
	return result, nil
}
