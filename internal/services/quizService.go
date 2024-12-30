package services

import (
	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/repository"
	"github.com/adon988/go_api_example/internal/responses"
	"gorm.io/gorm"
)

type QiuzServiceInterface interface {
	CreateQuiz(quiz models.Quiz) error
	CheckQuizExist(quiz_id string) bool
	InitQuizWithTx(quiz models.Quiz, quizAnswerRecord models.QuizAnswerRecord) error
	CreateQuizAnswerRecord(quizAnswerRecord models.QuizAnswerRecord) error
	UpdateQuizAnswerRecord(quizAnswerRecord models.QuizAnswerRecord) error
	GetQuizByMember(quiz_id string, member_id string) (responses.QuizWithAnswer, error)
	GetQuizsListWithAnswersByMember(member_id string, page int32) ([]responses.QuizWithAnswers, int32, error)
}

type QuizService struct {
	db               *gorm.DB
	quiz             repository.QuizRepository
	quizAnswerRecord repository.QuizAnswerRecordRepository
}

func NewQuizService(db *gorm.DB) QiuzServiceInterface {
	return &QuizService{
		db:               db,
		quiz:             repository.NewQuizRepository(db),
		quizAnswerRecord: repository.NewQuizAnswerRecordRepository(db),
	}
}

func (service QuizService) InitQuizWithTx(quiz models.Quiz, quizAnswerRecord models.QuizAnswerRecord) error {
	tx := service.db.Begin()
	NewTransactionQuizService := NewQuizService(tx)
	//grorm transaction
	err := NewTransactionQuizService.CreateQuiz(quiz)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = NewTransactionQuizService.CreateQuizAnswerRecord(quizAnswerRecord)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}

func (service QuizService) CreateQuiz(quiz models.Quiz) error {
	err := service.quiz.CreateQuiz(quiz)
	if err != nil {
		return err
	}

	return nil
}

func (service QuizService) CheckQuizExist(quiz_id string) bool {
	return service.quiz.CheckQuizExist(quiz_id)
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

func (service QuizService) GetQuizByMember(quiz_id string, member_id string) (responses.QuizWithAnswer, error) {
	quiz, err := service.quiz.GetQuizByMember(quiz_id, member_id)
	if err != nil {
		return responses.QuizWithAnswer{}, err
	}

	data := responses.QuizWithAnswer{
		QuizId:             quiz.QuizId,
		QuizAnswerRecordId: quiz.QuizAnswerRecordId,
		CreaterID:          quiz.CreaterID,
		QuestionType:       quiz.QuestionType,
		Topic:              quiz.Topic,
		Type:               quiz.Type,
		Info:               quiz.Info,
		Content:            quiz.Content,
		AnswerQuestion:     quiz.AnswerQuestion,
		Status:             quiz.Status,
		DueDate:            quiz.DueDate,
		FailedAnswerCount:  quiz.FailedAnswerCount,
		TotalQuestionCount: quiz.TotalQuestionCount,
		FailedLogs:         quiz.FailedLogs,
		Scope:              quiz.Scope,
	}

	return data, nil
}

func (service QuizService) GetQuizsListWithAnswersByMember(member_id string, page int32) ([]responses.QuizWithAnswers, int32, error) {
	quizzes, count, err := service.quiz.GetQuizsListWithAnswersByMember(member_id, page)
	if err != nil || count == 0 {
		return nil, 0, err
	}

	quizList := make([]responses.QuizWithAnswers, len(quizzes))
	for i, quiz := range quizzes {
		quizList[i] = responses.QuizWithAnswers{
			QuizId:             quiz.QuizId,
			QuizAnswerRecordId: quiz.QuizAnswerRecordId,
			CreaterID:          quiz.CreaterID,
			QuestionType:       quiz.QuestionType,
			Topic:              quiz.Topic,
			Type:               quiz.Type,
			Info:               quiz.Info,
			Status:             quiz.Status,
			DueDate:            quiz.DueDate,
			FailedAnswerCount:  quiz.FailedAnswerCount,
			TotalQuestionCount: quiz.TotalQuestionCount,
			FailedLogs:         quiz.FailedLogs,
			Scope:              quiz.Scope,
		}
	}
	return quizList, count, nil
}
