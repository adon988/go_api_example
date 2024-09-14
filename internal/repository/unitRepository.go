package repository

import (
	"fmt"

	"github.com/adon988/go_api_example/internal/models"
	"gorm.io/gorm"
)

type UnitRepository interface {
	CreateUnit(unit models.Unit) error
	UpdateUnit(unit models.Unit) error
	DeleteUnit(id string) error
	GetUnitByMemberID(member_id string) ([]models.Unit, error)
	GetUnitsByCourseID(course_id string) ([]models.Unit, error)
	GetUnitByMemberIDAndUnitID(member_id string, unit_id string) (models.Unit, error)
}

type UnitRepositoryImpl struct {
	DB *gorm.DB
}

func NewUnitRepository(db *gorm.DB) UnitRepository {
	return &UnitRepositoryImpl{DB: db}
}

func (r *UnitRepositoryImpl) CreateUnit(unit models.Unit) error {
	err := r.DB.Create(&unit)

	if err != nil {
		return err.Error
	}

	return nil
}

func (r *UnitRepositoryImpl) UpdateUnit(unit models.Unit) error {

	err := r.DB.Where("id = ?", unit.Id).Updates(unit)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return fmt.Errorf("no unit found with id: %s", unit.Id)
	}
	return nil
}

func (r *UnitRepositoryImpl) DeleteUnit(id string) error {
	result := r.DB.Delete(&models.Unit{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no unit found with id: %s", id)
	}

	return nil
}

func (r *UnitRepositoryImpl) GetUnitByMemberID(member_id string) ([]models.Unit, error) {
	var units []models.Unit
	err := r.DB.Model(&models.Unit{}).Joins("JOIN unit_permissions ON units.id = unit_permissions.entity_id").Where("unit_permissions.member_id = ?", member_id).Find(&units)
	if err.Error != nil {
		return nil, err.Error
	}

	return units, nil
}

func (r *UnitRepositoryImpl) GetUnitsByCourseID(course_id string) ([]models.Unit, error) {
	var units []models.Unit
	err := r.DB.Model(&models.Unit{}).Where("course_id = ?", course_id).Find(&units)
	if err.Error != nil {
		return nil, err.Error
	}

	return units, nil
}

func (r *UnitRepositoryImpl) GetUnitByMemberIDAndUnitID(member_id string, unit_id string) (models.Unit, error) {
	var unit models.Unit
	err := r.DB.Model(&models.Unit{}).Joins("JOIN unit_permissions ON units.id = unit_permissions.entity_id").Where("unit_permissions.member_id = ? AND units.id = ?", member_id, unit_id).Find(&unit)
	if err.Error != nil {
		return models.Unit{}, err.Error
	}
	return unit, nil
}
