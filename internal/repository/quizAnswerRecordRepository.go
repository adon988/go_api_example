package repository

import (
	"github.com/adon988/go_api_example/internal/models"
	"gorm.io/gorm"
)

type QuizAnswerRecordRepository interface {
	CreateQuizAnswerRecord(data models.QuizAnswerRecord) error
	UpdateQuizAnswerRecord(data models.QuizAnswerRecord) error
}

type quizAnswerRecordRepository struct {
	db *gorm.DB
}

func NewQuizAnswerRecordRepository(db *gorm.DB) QuizAnswerRecordRepository {
	return &quizAnswerRecordRepository{db: db}
}

func (r *quizAnswerRecordRepository) CreateQuizAnswerRecord(data models.QuizAnswerRecord) error {
	err := r.db.Create(&data)
	return err.Error
}

func (r *quizAnswerRecordRepository) UpdateQuizAnswerRecord(data models.QuizAnswerRecord) error {
	if err := r.db.Where("id = ?", data.Id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
