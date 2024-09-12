package services

import (
	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/repository"
	"github.com/adon988/go_api_example/internal/utils"
	"gorm.io/gorm"
)

type UnitServiceInterface interface {
	CreateUnitNPermission(member_id string, role int32, unit models.Unit) error
	GetUnit(member_id string) ([]models.Unit, error)
	UpdateUnit(unit models.Unit) error
	DeleteUnit(id string) error
	GetUnitPermissionByMemberIDAndUnitID(member_id string, unit_id string) (models.UnitPermission, error)
	AssignUnitPermission(unitPermission models.UnitPermission) error
	IsMemberWithEditorPermissionOnUnit(member_id string, unit_id string) (models.UnitPermission, error)
}

type UnitService struct {
	unit           repository.UnitRepository
	unitPermission repository.UnitPermissionRepository
}

func NewUnitService(db *gorm.DB) UnitServiceInterface {
	return &UnitService{
		unit:           repository.NewUnitRepository(db),
		unitPermission: repository.NewUnitPermissionRepository(db),
	}
}

func (service UnitService) IsMemberWithEditorPermissionOnUnit(member_id string, unit_id string) (models.UnitPermission, error) {
	unitPerm, err := service.GetUnitPermissionByMemberIDAndUnitID(member_id, unit_id)
	if err != nil {
		return unitPerm, err
	}

	err = utils.CheckRoleWithEditorPermission(unitPerm.Role)
	if err != nil {
		return unitPerm, err
	}
	return unitPerm, nil
}

func (service UnitService) CreateUnitNPermission(member_id string, role int32, unit models.Unit) error {
	unitPermission := models.UnitPermission{
		MemberId: member_id,
		EntityId: unit.Id,
		Role:     role,
	}

	result := service.unitPermission.CreateUnitPermission(unitPermission)
	if result != nil {
		return result
	}

	err := service.unit.CreateUnit(unit)
	if err != nil {
		return err
	}
	return nil
}

func (service UnitService) GetUnit(member_id string) ([]models.Unit, error) {
	return service.unit.GetUnitByMemberID(member_id)
}

func (service UnitService) UpdateUnit(unit models.Unit) error {
	return service.unit.UpdateUnit(unit)
}

func (service UnitService) DeleteUnit(id string) error {
	if err := service.unitPermission.DeleteUnitPermission(id); err != nil {
		return err
	}
	return service.unit.DeleteUnit(id)
}

func (service UnitService) GetUnitPermissionByMemberIDAndUnitID(member_id string, unit_id string) (models.UnitPermission, error) {
	return service.unitPermission.GetUnitPermissionByMemberIDAndUnitID(member_id, unit_id)
}

func (service UnitService) AssignUnitPermission(unitPermission models.UnitPermission) error {
	return service.unitPermission.AssignUnitPermission(unitPermission)
}
