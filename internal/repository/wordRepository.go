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
	GetWordByUnitID(unit_id string) ([]models.Word, error)
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

func (r *WordRepositoryImpl) GetWordByUnitID(unit_id string) ([]models.Word, error) {
	var words []models.Word

	err := r.DB.Where("unit_id = ?", unit_id).Find(&words)
	if err.Error != nil {
		return nil, err.Error
	}

	return words, nil
}
