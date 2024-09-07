package repository

import (
	"fmt"

	"github.com/adon988/go_api_example/internal/models"
	"gorm.io/gorm"
)

type UnitPermissionRepository interface {
	CreateUnitPermission(unitPermission models.UnitPermission) error
	GetUnitPermissionByMemberID(member_id string) ([]models.UnitPermission, error)
	GetUnitPermissionByMemberIDAndUnitID(member_id string, unit_id string) (models.UnitPermission, error)
	UpdateUnitPermission(unitPermission models.UnitPermission) error
	DeleteUnitPermission(id string) error
	AssignUnitPermission(unitPermission models.UnitPermission) error
}

func NewUnitPermissionRepository(db *gorm.DB) UnitPermissionRepository {
	return &UnitPermissionImpl{DB: db}
}

type UnitPermissionImpl struct {
	DB *gorm.DB
}

func (r *UnitPermissionImpl) CreateUnitPermission(unitPermission models.UnitPermission) error {
	hasRole := false
	for _, v := range models.PermissionRoleEnum {
		if v == unitPermission.Role {
			hasRole = true
			break
		}
	}
	if !hasRole {
		return fmt.Errorf("role not found")
	}
	result := r.DB.Create(&unitPermission)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UnitPermissionImpl) GetUnitPermissionByMemberID(member_id string) ([]models.UnitPermission, error) {
	var unitPerms []models.UnitPermission
	result := r.DB.Find(&unitPerms, "member_id = ?", member_id)
	if result.Error != nil {
		return nil, result.Error
	}
	return unitPerms, nil
}

func (r *UnitPermissionImpl) GetUnitPermissionByMemberIDAndUnitID(member_id string, unit_id string) (models.UnitPermission, error) {
	var unitPerm models.UnitPermission
	result := r.DB.Find(&unitPerm, "member_id = ? AND entity_id = ?", member_id, unit_id)
	if result.Error != nil {
		return unitPerm, result.Error
	}
	if result.RowsAffected == 0 {
		return unitPerm, gorm.ErrRecordNotFound
	}
	return unitPerm, nil
}

func (r *UnitPermissionImpl) UpdateUnitPermission(unitPermission models.UnitPermission) error {
	result := r.DB.Updates(&unitPermission)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UnitPermissionImpl) DeleteUnitPermission(id string) error {
	result := r.DB.Delete(&models.UnitPermission{}, "entity_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *UnitPermissionImpl) AssignUnitPermission(unitPermission models.UnitPermission) error {
	hasRole := false
	for _, v := range models.PermissionRoleEnum {
		if v == unitPermission.Role {
			hasRole = true
			break
		}
	}
	if !hasRole {
		return fmt.Errorf("role not found")
	}

	result := r.DB.Model(&models.UnitPermission{}).Where("member_id = ? AND entity_id = ?", unitPermission.MemberId, unitPermission.EntityId).Updates(&unitPermission)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		result = r.DB.Create(&unitPermission)
	}

	return result.Error
}
