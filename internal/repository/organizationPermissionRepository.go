package repository

import (
	"fmt"

	"github.com/adon988/go_api_example/internal/models"
	"gorm.io/gorm"
)

type OrganizationPermissionRepository interface {
	CreateOrganizationPermission(organization_permissions models.OrganizationPermission) error
	GetOrganizationPermissionByMemberID(member_id string) ([]models.OrganizationPermission, error)
	UpdateOrganizationPermission(organization_permissions models.OrganizationPermission) error
	DeleteOrganizationPermission(id string) error
}

func NewOrganizationPermission(db *gorm.DB) OrganizationPermissionRepository {
	return &OrganizationPermissionImpl{DB: db}
}

type OrganizationPermissionImpl struct {
	DB *gorm.DB
}

func (r *OrganizationPermissionImpl) CreateOrganizationPermission(organization_permissions models.OrganizationPermission) error {
	if _, ok := models.PermissionRoleEnum[organization_permissions.Role]; !ok {
		return fmt.Errorf("role %s not found", organization_permissions.Role)
	}

	result := r.DB.Create(&organization_permissions)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *OrganizationPermissionImpl) GetOrganizationPermissionByMemberID(member_id string) ([]models.OrganizationPermission, error) {
	var orgPerms []models.OrganizationPermission
	result := r.DB.Find(&orgPerms, "member_id = ?", member_id)
	if result.Error != nil {
		return nil, result.Error
	}
	return orgPerms, nil
}

func (r *OrganizationPermissionImpl) UpdateOrganizationPermission(organization_permissions models.OrganizationPermission) error {
	result := r.DB.Save(&organization_permissions)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *OrganizationPermissionImpl) DeleteOrganizationPermission(id string) error {
	result := r.DB.Delete(&models.OrganizationPermission{}, "entity_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
