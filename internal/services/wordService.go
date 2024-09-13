package services

import (
	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/repository"
	"gorm.io/gorm"
)

type WordServiceInterface interface {
	CreateWord(word models.Word) error
	UpdateWord(word models.Word) error
	DeleteWord(id string) error
	GetWordByUnitID(unit_id string) ([]models.Word, error)
}

type WordService struct {
	word repository.WordRepository
}

func NewWordService(db *gorm.DB) WordServiceInterface {
	return &WordService{
		word: repository.NewWordRepositoryImpl(db),
	}
}

func (service WordService) CreateWord(word models.Word) error {
	err := service.word.CreateWord(word)
	if err != nil {
		return err
	}

	return nil
}

func (service WordService) UpdateWord(word models.Word) error {
	err := service.word.UpdateWord(word)
	if err != nil {
		return err
	}

	return nil
}

func (service WordService) DeleteWord(id string) error {
	err := service.word.DeleteWord(id)
	if err != nil {
		return err
	}

	return nil
}

func (service WordService) GetWordByUnitID(unit_id string) ([]models.Word, error) {
	words, err := service.word.GetWordByUnitID(unit_id)
	if err != nil {
		return nil, err
	}

	return words, nil
}
