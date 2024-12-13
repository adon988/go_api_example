package services

import (
	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/repository"
	"gorm.io/gorm"
)

type WordServiceInterface interface {
	CreateWord(word models.Word) error
	UpdateWord(word models.Word) error
	DeleteWord(word_id string) error
	GetWordByMemberIDAndUnitID(member_id string, unit_id string) ([]models.Word, error)
	GetWordByMemberIDAndCourseID(member_id string, course_id string) ([]models.Word, error)
	CheckWordPermissionByMemberIDAndWordID(member_id string, word_id string) (models.Word, error)
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
	result := service.word.CreateWord(word)
	if result != nil {
		return result
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

func (service WordService) DeleteWord(word_id string) error {
	err := service.word.DeleteWord(word_id)
	if err != nil {
		return err
	}

	return nil
}

func (service WordService) GetWordByMemberIDAndUnitID(member_id string, unit_id string) ([]models.Word, error) {
	var words []models.Word
	words, err := service.word.GetWordByMemberIDAndUnitID(member_id, unit_id)
	if err != nil {
		return []models.Word{}, err
	}

	return words, nil
}

func (service WordService) CheckWordPermissionByMemberIDAndWordID(member_id string, word_id string) (models.Word, error) {
	words, err := service.word.CheckWordPermissionByMemberIDAndWordID(member_id, word_id)
	if err != nil {
		return models.Word{}, err
	}

	return words, nil
}

func (service WordService) GetWordByMemberIDAndCourseID(member_id string, course_id string) ([]models.Word, error) {
	words, err := service.word.GetWordByMemberIDAndCourseID(member_id, course_id)
	if err != nil {
		return nil, err
	}

	return words, nil
}
