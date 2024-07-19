package services

import (
	"github.com/adon988/go_api_example/api/repository"
	"github.com/adon988/go_api_example/models"
	"gorm.io/gorm"
)

type RoleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{
		roleRepo: repository.NewRoleRepository(db),
	}
}

func (service RoleService) GetRole() ([]*models.Role, error) {
	return service.roleRepo.GetRole()
}

func (service RoleService) CreateRole(role models.Role) error {
	return service.roleRepo.CreateRole(role)
}

func (service RoleService) UpdateRole(role models.Role) error {
	return service.roleRepo.UpdateRole(role)
}

func (service RoleService) DeleteRole(id int32) error {
	return service.roleRepo.DeleteRole(id)
}
