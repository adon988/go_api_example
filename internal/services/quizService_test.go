package services

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestInitQuiz(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Quiz{}, &models.QuizAnswerRecord{})

	quiz := models.Quiz{
		Id:           "1",
		CreaterId:    "1",
		QuestionType: "multiple_choice",
		Topic:        1,
		Type:         1,
		Info:         nil,
		Content:      nil,
	}

	quizAnswerRecord := models.QuizAnswerRecord{
		Id:                 "1",
		QuizId:             "1",
		MemberId:           "1",
		AnswerQuestion:     &json.RawMessage{},
		Status:             1,
		DueDate:            func() *time.Time { t := time.Now().AddDate(0, 0, 30); return &t }(),
		FailedAnswerCount:  func(i int32) *int32 { return &i }(10),
		TotalQuestionCount: func(i int32) *int32 { return &i }(20),
		FailedLogs:         &json.RawMessage{},
		Scope:              func(i int32) *int32 { return &i }(50),
	}
	service := NewQuizService(mockDB)
	err := service.InitQuizWithTx(quiz, quizAnswerRecord)
	assert.Nil(t, err)

}

func TestQuizService_CreateQuiz(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Quiz{})
	service := NewQuizService(mockDB)
	quiz := models.Quiz{
		Id:           "1",
		CreaterId:    "1",
		QuestionType: "multiple_choice",
		Topic:        1,
		Type:         1,
		Info:         nil,
		Content:      nil,
	}

	err := service.CreateQuiz(quiz)
	assert.Nil(t, err)

	// 檢查 quiz 是否被創建
	var quizzes []models.Quiz
	mockDB.Find(&quizzes)
	assert.Equal(t, 1, len(quizzes))
	assert.Equal(t, quiz.Id, quizzes[0].Id)

}
func TestQuizService_CreateQuizAnswerRecord(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.QuizAnswerRecord{})
	service := NewQuizService(mockDB)
	quizAnswerRecord := models.QuizAnswerRecord{
		Id:                 "1",
		QuizId:             "1",
		MemberId:           "1",
		AnswerQuestion:     &json.RawMessage{},
		Status:             1,
		DueDate:            func() *time.Time { t := time.Now().AddDate(0, 0, 30); return &t }(),
		FailedAnswerCount:  func(i int32) *int32 { return &i }(10),
		TotalQuestionCount: func(i int32) *int32 { return &i }(20),
		FailedLogs:         &json.RawMessage{},
		Scope:              func(i int32) *int32 { return &i }(50),
	}
	err := service.CreateQuizAnswerRecord(quizAnswerRecord)
	assert.Nil(t, err)

	// 檢查 quizAnswerRecord 是否被創建
	var quizAnswerRecords []models.QuizAnswerRecord
	mockDB.Find(&quizAnswerRecords)
	assert.Equal(t, 1, len(quizAnswerRecords))
	assert.Equal(t, quizAnswerRecord.Id, quizAnswerRecords[0].Id)
}
func TestQuizService_UpdateQuizAnswerRecord(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.QuizAnswerRecord{})
	service := NewQuizService(mockDB)
	quizAnswerRecord := models.QuizAnswerRecord{
		Id:                 "1",
		QuizId:             "1",
		MemberId:           "1",
		AnswerQuestion:     &json.RawMessage{},
		Status:             1,
		DueDate:            func() *time.Time { t := time.Now().AddDate(0, 0, 30); return &t }(),
		FailedAnswerCount:  func(i int32) *int32 { return &i }(10),
		TotalQuestionCount: func(i int32) *int32 { return &i }(20),
		FailedLogs:         &json.RawMessage{},
		Scope:              func(i int32) *int32 { return &i }(50),
	}
	err := mockDB.Create(&quizAnswerRecord)
	assert.Nil(t, err.Error)

	quizAnswerRecordUpdate := models.QuizAnswerRecord{
		Id:     quizAnswerRecord.Id,
		Status: 2,
	}

	service.UpdateQuizAnswerRecord(quizAnswerRecordUpdate)
	var quizAnswerRecords []models.QuizAnswerRecord
	mockDB.Find(&quizAnswerRecords)
	assert.Equal(t, 1, len(quizAnswerRecords))
	assert.Equal(t, quizAnswerRecordUpdate.Status, quizAnswerRecords[0].Status)

}
func TestQuizService_GetQuizByMember(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Quiz{}, &models.QuizAnswerRecord{})
	service := NewQuizService(mockDB)
	quiz := models.Quiz{
		Id:           "1",
		CreaterId:    "1",
		QuestionType: "multiple_choice",
		Topic:        1,
		Type:         1,
		Info:         &json.RawMessage{},
		Content:      &json.RawMessage{},
	}

	err := service.CreateQuiz(quiz)
	assert.Nil(t, err)

	quizAnswerRecord := models.QuizAnswerRecord{
		Id:                 "1",
		QuizId:             "1",
		MemberId:           "1",
		AnswerQuestion:     &json.RawMessage{},
		Status:             1,
		DueDate:            func() *time.Time { t := time.Now().AddDate(0, 0, 30); return &t }(),
		FailedAnswerCount:  func(i int32) *int32 { return &i }(10),
		TotalQuestionCount: func(i int32) *int32 { return &i }(20),
		FailedLogs:         &json.RawMessage{},
		Scope:              func(i int32) *int32 { return &i }(50),
	}

	err = service.CreateQuizAnswerRecord(quizAnswerRecord)
	assert.Nil(t, err)

	quizWithAnswer, errors := service.GetQuizByMember(quiz.Id, quizAnswerRecord.MemberId)
	assert.Nil(t, errors)
	assert.Equal(t, quizWithAnswer.QuizId, quizAnswerRecord.QuizId)

}
func TestQuizService_GetQuizsListWithAnswersByMember(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Quiz{}, &models.QuizAnswerRecord{})
	service := NewQuizService(mockDB)
	quiz := models.Quiz{
		Id:           "1",
		CreaterId:    "1",
		QuestionType: "multiple_choice",
		Topic:        1,
		Type:         1,
		Info:         &json.RawMessage{},
		Content:      &json.RawMessage{},
	}

	err := service.CreateQuiz(quiz)
	assert.Nil(t, err)

	quizAnswerRecord := models.QuizAnswerRecord{
		Id:                 "1",
		QuizId:             "1",
		MemberId:           "1",
		AnswerQuestion:     &json.RawMessage{},
		Status:             1,
		DueDate:            func() *time.Time { t := time.Now().AddDate(0, 0, 30); return &t }(),
		FailedAnswerCount:  func(i int32) *int32 { return &i }(10),
		TotalQuestionCount: func(i int32) *int32 { return &i }(20),
		FailedLogs:         &json.RawMessage{},
		Scope:              func(i int32) *int32 { return &i }(50),
	}

	err = service.CreateQuizAnswerRecord(quizAnswerRecord)
	assert.Nil(t, err)

	QuizWithAnswer, count, errors := service.GetQuizsListWithAnswersByMember(quizAnswerRecord.MemberId, 0)
	assert.Nil(t, errors)
	assert.Equal(t, int32(1), count)
	assert.Equal(t, 1, len(QuizWithAnswer))
	assert.Equal(t, QuizWithAnswer[0].QuizId, quizAnswerRecord.QuizId)
}
