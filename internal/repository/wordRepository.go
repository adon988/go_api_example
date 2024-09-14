package repository

import (
	"fmt"

	"github.com/adon988/go_api_example/internal/models"
	"gorm.io/gorm"
)

type WordRepository interface {
	CreateWord(word models.Word) error
	UpdateWord(word models.Word) error
	DeleteWord(id string) error
	GetWordByMemberIDAndUnitID(member_id string, unit_id string) ([]models.Word, error)
	GetWordByMemberIDAndCourseID(member_id string, couse_id string) ([]models.Word, error)
	CheckWordPermissionByMemberIDAndWordID(member_id string, word_id string) (models.Word, error)
}

type WordRepositoryImpl struct {
	DB *gorm.DB
}

func NewWordRepositoryImpl(db *gorm.DB) WordRepository {
	return &WordRepositoryImpl{DB: db}
}

func (r *WordRepositoryImpl) CreateWord(word models.Word) error {
	err := r.DB.Create(&word)

	if err != nil {
		return err.Error
	}

	return nil
}

func (r *WordRepositoryImpl) UpdateWord(word models.Word) error {

	err := r.DB.Where("id = ?", word.Id).Updates(word)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return fmt.Errorf("no word found with id: %s", word.Id)
	}
	return nil
}

func (r *WordRepositoryImpl) DeleteWord(id string) error {
	result := r.DB.Delete(&models.Word{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no word found with id: %s", id)
	}

	return nil
}

func (r *WordRepositoryImpl) GetWordByMemberIDAndUnitID(member_id string, unit_id string) ([]models.Word, error) {
	var words []models.Word

	err := r.DB.Model(&models.Word{}).Joins("JOIN unit_permissions ON words.unit_id = unit_permissions.entity_id").Where("unit_permissions.member_id = ? AND words.unit_id = ?", member_id, unit_id).Find(&words)
	if err.Error != nil {
		return nil, err.Error
	}

	return words, nil
}

func (r *WordRepositoryImpl) GetWordByMemberIDAndCourseID(member_id string, course_id string) ([]models.Word, error) {
	var words []models.Word

	err := r.DB.Model(&models.Word{}).Joins("JOIN unit_permissions ON words.unit_id = unit_permissions.entity_id").Joins("JOIN units ON words.unit_id = units.id").Joins("JOIN course_permissions ON units.course_id = course_permissions.entity_id").Where("unit_permissions.member_id = ? AND course_permissions.member_id = ? AND course_permissions.entity_id = ?", member_id, member_id, course_id).Find(&words)
	if err.Error != nil {
		return nil, err.Error
	}

	return words, nil

}

func (r *WordRepositoryImpl) CheckWordPermissionByMemberIDAndWordID(member_id string, word_id string) (models.Word, error) {
	var word models.Word
	err := r.DB.Model(&models.Word{}).Joins("JOIN unit_permissions ON words.unit_id = unit_permissions.entity_id").Where("words.id = ?", word_id).Find(&word)
	if err.Error != nil {
		return models.Word{}, err.Error
	}

	return word, nil
}
