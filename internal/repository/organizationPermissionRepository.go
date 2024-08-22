package repository

import (
	"fmt"

	"github.com/adon988/go_api_example/internal/models"
	"gorm.io/gorm"
)

type OrganizationPermissionRepository interface {
	CreateOrganizationPermission(organization_permissions models.OrganizationPermission) error
	GetOrganizationPermissionByMemberID(member_id string) ([]models.OrganizationPermission, error)
	GetOrganizationPermissionByOrganizationIDAndMemberID(member_id string, organization_id string) (models.OrganizationPermission, error)
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

	hasRole := false
	for _, v := range models.PermissionRoleEnum {
		if v == organization_permissions.Role {
			hasRole = true
			break
		}
	}
	if !hasRole {
		return fmt.Errorf("role not found")
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

func (r *OrganizationPermissionImpl) GetOrganizationPermissionByOrganizationIDAndMemberID(member_id string, organization_id string) (models.OrganizationPermission, error) {
	var orgPerm models.OrganizationPermission
	result := r.DB.Find(&orgPerm, "member_id = ? AND entity_id = ?", member_id, organization_id)
	if result.Error != nil {
		return orgPerm, result.Error
	}
	if result.RowsAffected == 0 {
		return orgPerm, gorm.ErrRecordNotFound
	}
	return orgPerm, nil
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
