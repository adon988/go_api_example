package repository

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestQuizRepository_CreateQuiz(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Quiz{}, &models.QuizAnswerRecord{})

	repo := NewQuizRepository(mockDB)

	quiz := models.Quiz{
		Id:           "1",
		CreaterId:    "1",
		QuestionType: "multiple_choice, true_false, full_in_blank",
		Topic:        1,
		Type:         1,
		Info:         nil,
		Content:      nil,
	}
	err := repo.CreateQuiz(quiz)
	assert.Nil(t, err)

	// 檢查 quiz 是否被創建
	var quizzes []models.Quiz
	mockDB.Find(&quizzes)
	assert.Equal(t, 1, len(quizzes))
	assert.Equal(t, quiz.Id, quizzes[0].Id)

	fmt.Println("Result:")
	fmt.Println(quizzes)
}

func TestQuizRepository_CheckQuizExist(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Quiz{}, &models.QuizAnswerRecord{})

	repo := NewQuizRepository(mockDB)

	quiz := models.Quiz{
		Id:           "1",
		CreaterId:    "1",
		QuestionType: "multiple_choice, true_false, full_in_blank",
		Topic:        1,
		Type:         1,
		Info:         nil,
		Content:      nil,
	}
	err := repo.CreateQuiz(quiz)
	assert.Nil(t, err)

	// 檢查 quiz 是否存在
	exist := repo.CheckQuizExist(quiz.Id)
	assert.Equal(t, true, exist)

	fmt.Println("Result:")
	fmt.Println(exist)
}

func TestQuizRepository_GetQuizByMember(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Quiz{}, &models.QuizAnswerRecord{})

	repo := NewQuizRepository(mockDB)
	info := models.QuizInfo{
		QuizCount: 1,
		RetryTime: 1,
		Organization: models.ClassInfo{
			Id:    "1",
			Title: "Organization 1",
		},
	}
	infoRaw := utils.MarshalJSONToRaw(info)

	contentItems := models.ContentItems{}
	ContentItem := models.ContentItem{
		QuestionType: "multiple_choice",
		Question: []models.Title{
			{Title: "What is the capital of France?", Id: "1"},
			{Title: "What is the capital of France2?", Id: "2"},
			{Title: "What is the capital of France3?", Id: "3"},
		},
		Answer:        "0", //word_id
		Word:          "capital",
		Definition:    "The city or town that functions as the seat of government and administrative centre of a country or region.",
		Pronunciation: "ˈkapɪt(ə)l",
	}
	contentItems.ContentItems = append(contentItems.ContentItems, ContentItem)
	ContentItem = models.ContentItem{
		QuestionType: "true_false",
		Question: []models.Title{
			{Title: "飛機", Id: "1"},
		},
		Answer:        "0",
		WordId:        "2",
		Word:          "Application",
		Definition:    "This appliction will help to improve your efficiency.",
		Pronunciation: "ˈkapɪt(ə)l",
	}
	contentItems.ContentItems = append(contentItems.ContentItems, ContentItem)
	contentItem := models.ContentItem{
		QuestionType: "full_in_blank",
		Question: []models.Title{
			{Title: "App[*]lca[*]ion"},
		},
		Answer:        "i,t",
		WordId:        "2",
		Word:          "Application",
		Definition:    "This appliction will help to improve your efficiency.",
		Pronunciation: "ˈkapɪt(ə)l",
	}
	contentItems.ContentItems = append(contentItems.ContentItems, contentItem)

	contentRaw := utils.MarshalJSONToRaw(contentItems)
	quiz := models.Quiz{
		Id:           "1",
		CreaterId:    "1",
		QuestionType: "multiple_choice",
		Topic:        1,
		Type:         1,
		Info:         infoRaw,
		Content:      contentRaw,
	}
	err := repo.CreateQuiz(quiz)
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
	repoQuizAnswerRecord := NewQuizAnswerRecordRepository(mockDB)
	err = repoQuizAnswerRecord.CreateQuizAnswerRecord(quizAnswerRecord)
	assert.Nil(t, err)

	QuizWithAnswer, errs := repo.GetQuizByMember(quiz.Id, quizAnswerRecord.MemberId)
	assert.Nil(t, errs)
	assert.Equal(t, QuizWithAnswer.QuizId, "1")
	assert.Equal(t, quizAnswerRecord.Id, QuizWithAnswer.QuizAnswerRecordId)
	quizInfo := &models.QuizInfo{}
	json.Unmarshal([]byte(QuizWithAnswer.Info), quizInfo)
	assert.Equal(t, quizInfo.QuizCount, info.QuizCount)
	assert.Equal(t, quizInfo.Organization.Title, info.Organization.Title)
	quizContent := &models.ContentItems{}
	json.Unmarshal([]byte(QuizWithAnswer.Content), &quizContent)
	assert.Equal(t, quizContent.ContentItems[0].QuestionType, contentItems.ContentItems[0].QuestionType)
	assert.Equal(t, quizContent.ContentItems[1].Word, contentItems.ContentItems[1].Word)
	assert.Equal(t, quizContent.ContentItems[2].Word, contentItems.ContentItems[2].Word)
	fmt.Println("Result:")
	fmt.Println(quiz)
}

func TestQuizRepository_GetQuizsListWithAnswersByMember(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Quiz{}, &models.QuizAnswerRecord{})

	repo := NewQuizRepository(mockDB)

	quiz := models.Quiz{
		Id:           "1",
		CreaterId:    "1",
		QuestionType: "multiple_choice, true_false, full_in_blank",
		Topic:        1,
		Type:         1,
		Info:         &json.RawMessage{},
		Content:      &json.RawMessage{},
	}
	err := repo.CreateQuiz(quiz)
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
	repoQuizAnswerRecord := NewQuizAnswerRecordRepository(mockDB)
	err = repoQuizAnswerRecord.CreateQuizAnswerRecord(quizAnswerRecord)
	assert.Nil(t, err)

	quizzes, count, err := repo.GetQuizsListWithAnswersByMember(quizAnswerRecord.MemberId, 0)
	assert.Nil(t, err)
	assert.Equal(t, int32(1), count)
	assert.Equal(t, 1, len(quizzes))
	assert.Equal(t, quiz.Id, quizzes[0].QuizId)
	assert.Equal(t, quizAnswerRecord.Id, quizzes[0].QuizAnswerRecordId)

	fmt.Println("Result:")
	fmt.Println(quizzes)
}
