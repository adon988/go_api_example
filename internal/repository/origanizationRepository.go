package repository

import (
	"fmt"

	"github.com/adon988/go_api_example/internal/models"
	"gorm.io/gorm"
)

type OriganizationRepository interface {
	CreateOriganization(origanization models.Organization) error
	DeleteOriganization(id string) error
	GetOriganizationByMemberID(id string) ([]models.Organization, error)
}

type OriganizationRepositoryImpl struct {
	DB          *gorm.DB
	SupportJSON bool
}

func NewOriganizationRepository(db *gorm.DB, support_json bool) OriganizationRepository {
	return &OriganizationRepositoryImpl{DB: db, SupportJSON: support_json}
}

// Implement the interface methods
var origanization []models.Organization

func (r *OriganizationRepositoryImpl) CreateOriganization(origanization models.Organization) error {
	err := r.DB.Create(&origanization)

	return err.Error
}
func (r *OriganizationRepositoryImpl) DeleteOriganization(id string) error {
	result := r.DB.Delete(origanization, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no origanization found with id: %s", id)
	}
	return nil
}

func (r *OriganizationRepositoryImpl) GetOriganizationByMemberID(id string) ([]models.Organization, error) {
	permissionQuery := fmt.Sprintf(`[{"user_id": "%s"}]`, id)
	query := "JSON_CONTAINS(permissions, ?)"

	if !r.SupportJSON {
		permissionQuery = "%\"user_id\":\"" + id + "\"%"
		query = "permissions LIKE ?"
	}

	if err := r.DB.Find(&origanization, query, permissionQuery).Error; err != nil {
		return nil, err
	}

	return origanization, nil
}
