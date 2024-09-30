package services

import (
	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/repository"
	"gorm.io/gorm"
)

type QiuzServiceInterface interface {
	CreateQuiz(quiz models.Quiz) error
	CreateQuizAnswerRecord(quizAnswerRecord models.QuizAnswerRecord) error
	UpdateQuizAnswerRecord(quizAnswerRecord models.QuizAnswerRecord) error
	GetQuizByMember(quiz_id string, member_id string) (models.QuizWithAnswer, error)
	GetQuizsListWithAnswersByMember(member_id string, page int32) ([]models.QuizWithAnswers, int32, error)
}

type QuizService struct {
	quiz             repository.QuizRepository
	quizAnswerRecord repository.QuizAnswerRecordRepository
}

func NewQuizService(db *gorm.DB) QiuzServiceInterface {
	return &QuizService{
		quiz:             repository.NewQuizRepository(db),
		quizAnswerRecord: repository.NewQuizAnswerRecordRepository(db),
	}
}

func (service QuizService) CreateQuiz(quiz models.Quiz) error {
	err := service.quiz.CreateQuiz(quiz)
	if err != nil {
		return err
	}

	return nil
}

func (service QuizService) CreateQuizAnswerRecord(quizAnswerRecord models.QuizAnswerRecord) error {
	err := service.quizAnswerRecord.CreateQuizAnswerRecord(quizAnswerRecord)
	if err != nil {
		return err
	}

	return nil
}

func (service QuizService) UpdateQuizAnswerRecord(quizAnswerRecord models.QuizAnswerRecord) error {
	err := service.quizAnswerRecord.UpdateQuizAnswerRecord(quizAnswerRecord)
	if err != nil {
		return err
	}

	return nil
}

func (service QuizService) GetQuizByMember(quiz_id string, member_id string) (models.QuizWithAnswer, error) {
	quiz, err := service.quiz.GetQuizByMember(quiz_id, member_id)
	if err != nil {
		return models.QuizWithAnswer{}, err
	}

	return quiz, nil
}

func (service QuizService) GetQuizsListWithAnswersByMember(member_id string, page int32) ([]models.QuizWithAnswers, int32, error) {
	quizzes, count, err := service.quiz.GetQuizsListWithAnswersByMember(member_id, page)
	if err != nil {
		return nil, 0, err
	}

	return quizzes, count, nil
}
