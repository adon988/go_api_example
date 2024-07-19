package repository

import (
	"fmt"

	"github.com/adon988/go_api_example/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRole() ([]*models.Role, error)
	CreateRole(role models.Role) error
	UpdateRole(role models.Role) error
	DeleteRole(id int32) error
}

type RoleRepositoryImpl struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{DB: db}
}

var role models.Role

func (r *RoleRepositoryImpl) GetRole() ([]*models.Role, error) {
	var roles []*models.Role
	if err := r.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleRepositoryImpl) CreateRole(role models.Role) error {
	err := r.DB.Create(&role)
	return err.Error
}

func (r *RoleRepositoryImpl) UpdateRole(role models.Role) error {
	if err := r.DB.Where("id = ?", role.Id).Updates(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *RoleRepositoryImpl) DeleteRole(id int32) error {
	result := r.DB.Delete(&role, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no role found with id: %d", id)
	}

	return nil
}
