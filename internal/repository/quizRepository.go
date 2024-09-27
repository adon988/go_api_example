package repository

import (
	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/utils"
	"gorm.io/gorm"
)

type QuizRepository interface {
	CreateQuiz(data models.Quiz) error
	GetQuizByMember(quiz_id string, member_id string) (models.QuizWithAnswer, error)
	GetQuizsListWithAnswersByMember(member_id string, page int32) ([]models.QuizWithAnswers, int32, error)
}

type quizRepositoryImpl struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) QuizRepository {
	return &quizRepositoryImpl{db: db}
}

func (r *quizRepositoryImpl) CreateQuiz(data models.Quiz) error {
	return r.db.Create(&data).Error
}

func (r *quizRepositoryImpl) GetQuizByMember(quiz_id string, member_id string) (models.QuizWithAnswer, error) {
	var quiz models.QuizWithAnswer
	err := r.db.Table("quizzes").
		Select("quizzes.id as quiz_id, quiz_answer_records.id as quiz_answer_record_id, quizzes.creater_id, quizzes.question_type, quizzes.topic, quizzes.type, quizzes.info, quizzes.content, quiz_answer_records.answer_question, quiz_answer_records.status, quiz_answer_records.due_date, quiz_answer_records.correct_answer_count, quiz_answer_records.total_question_count, quiz_answer_records.failed_logs, quiz_answer_records.scope").
		Joins("left join quiz_answer_records on quizzes.id = quiz_answer_records.quiz_id").
		Where(
			"quizzes.id = ? AND quiz_answer_records.member_id = ?",
			quiz_id, member_id,
		).
		Limit(1).
		Find(&quiz).Error

	return quiz, err
}

func (r *quizRepositoryImpl) GetQuizsListWithAnswersByMember(member_id string, page int32) ([]models.QuizWithAnswers, int32, error) {

	var quizzes []models.QuizWithAnswers
	err := r.db.Table("quizzes").
		Select("quizzes.id as quiz_id, quiz_answer_records.id as quiz_answer_record_id, quizzes.creater_id, quizzes.question_type, quizzes.topic, quizzes.type, quizzes.info, quiz_answer_records.status, quiz_answer_records.due_date, quiz_answer_records.correct_answer_count, quiz_answer_records.total_question_count, quiz_answer_records.failed_logs, quiz_answer_records.scope").
		Joins("left join quiz_answer_records on quizzes.id = quiz_answer_records.quiz_id").
		Where("quiz_answer_records.member_id = ?", member_id).
		Offset(int(page) * utils.Configs.Service.Page_Items).
		Limit(10).
		Find(&quizzes).Error

	if err != nil {
		return nil, 0, err
	}

	var count int64

	err = r.db.Table("quizzes").
		Select("quizzes.*, quiz_answer_records.*").
		Joins("left join quiz_answer_records on quizzes.id = quiz_answer_records.quiz_id").
		Where("quiz_answer_records.member_id = ?", member_id).
		Count(&count).Error

	return quizzes, int32(count), err
}
