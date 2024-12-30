package repository

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestQuizAnswerRecordRepository_CreateQuizAnswerRecord(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.QuizAnswerRecord{})

	repo := NewQuizAnswerRecordRepository(mockDB)

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
	err := repo.CreateQuizAnswerRecord(quizAnswerRecord)
	assert.Nil(t, err)

	// 檢查 quizAnswerRecord 是否被創建
	var quizAnswerRecords []models.QuizAnswerRecord
	mockDB.Find(&quizAnswerRecords)
	assert.Equal(t, 1, len(quizAnswerRecords))
	assert.Equal(t, quizAnswerRecord.Id, quizAnswerRecords[0].Id)

	fmt.Println("Result:")
	fmt.Println(quizAnswerRecords)
}

func TestQuizAnswerRecordRepository_UpdateQuizAnswerRecord(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.QuizAnswerRecord{})

	repo := NewQuizAnswerRecordRepository(mockDB)

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
	repo.UpdateQuizAnswerRecord(quizAnswerRecordUpdate)

	// 檢查 quizAnswerRecord 是否被更新
	var quizAnswerRecords []models.QuizAnswerRecord
	mockDB.Find(&quizAnswerRecords)
	assert.Equal(t, 1, len(quizAnswerRecords))
	assert.Equal(t, quizAnswerRecordUpdate.Id, quizAnswerRecords[0].Id)
	assert.Equal(t, quizAnswerRecordUpdate.Status, quizAnswerRecords[0].Status)

	fmt.Println("Result:")
	fmt.Println(quizAnswerRecordUpdate)
}
