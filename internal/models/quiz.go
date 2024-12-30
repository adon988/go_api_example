package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Quiz struct {
	Id             string           `gorm:"primaryKey;size:255"`
	CreaterId      string           `gorm:"size:24;index;comment:'only note the creater id, not the permission'"`
	QuestionType   string           `gorm:"size:255; comment:'type: multiple_choice, true_false, fill_in_the_blank'"`
	Topic          int32            `gorm:"type:tinyint(2); comment:'1:source, 2:target';default:1"`
	Type           int32            `gorm:"type:tinyint(2); comment:'1:quiz, 2:challenge';default:1"`
	OrganizationId *string          `gorm:"size:24"`
	CourseId       *string          `gorm:"size:24"`
	UnitId         *string          `gorm:"size:24"`
	Info           *json.RawMessage `gorm:"type:json;comment:'quiz info: exam date, retry times, org, course, unit...info'"`
	Content        *json.RawMessage `gorm:"type:json;comment:'quiz and anser info'"`
	CreatedAt      time.Time        // Automatically managed by GORM for creation time
	UpdatedAt      time.Time        // Automatically managed by GORM for update time
	DeletedAt      gorm.DeletedAt   `gorm:"index"`
}

// 自定義結構體，用於查詢帶有答案的問卷
type QuizWithAnswer struct {
	QuizId             string    `gorm:"column:quiz_id"`
	QuizAnswerRecordId string    `gorm:"column:quiz_answer_record_id"`
	CreaterID          string    `gorm:"column:creater_id"`
	QuestionType       string    `gorm:"column:question_type"`
	Topic              int32     `gorm:"column:topic"`
	Type               int32     `gorm:"column:type"`
	Info               string    `gorm:"column:info"`
	Content            string    `gorm:"column:content"`
	AnswerQuestion     string    `gorm:"column:answer_question"`
	Status             int32     `gorm:"column:status"`
	DueDate            time.Time `gorm:"column:due_date"`
	FailedAnswerCount  int32     `gorm:"column:failed_answer_count"`
	TotalQuestionCount int32     `gorm:"column:total_question_count"`
	FailedLogs         string    `gorm:"column:failed_logs"`
	Scope              int32     `gorm:"column:scope"`
}

type QuizWithAnswers struct {
	QuizId             string    `gorm:"column:quiz_id"`
	QuizAnswerRecordId string    `gorm:"column:quiz_answer_record_id"`
	CreaterID          string    `gorm:"column:creater_id"`
	QuestionType       string    `gorm:"column:question_type"`
	Topic              int32     `gorm:"column:topic"`
	Type               int32     `gorm:"column:type"`
	Info               string    `gorm:"column:info"`
	Status             int32     `gorm:"column:status"`
	DueDate            time.Time `gorm:"column:due_date"`
	FailedAnswerCount  int32     `gorm:"column:failed_answer_count"`
	TotalQuestionCount int32     `gorm:"column:total_question_count"`
	FailedLogs         string    `gorm:"column:failed_logs"`
	Scope              int32     `gorm:"column:scope"`
}

type ClassInfo struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}
type QuizInfo struct {
	QuizCount    int32     `json:"quiz_count"`
	QuizDays     int32     `json:"quiz_days"`
	RetryTime    int32     `json:"retry_time"`
	Organization ClassInfo `json:"organization"`
	Course       ClassInfo `json:"course"`
	Unit         ClassInfo `json:"unit"`
}
type Titles struct {
	Titles []Title `json:"titles"`
}

type Title struct {
	Title string `json:"title"`
	Id    string `json:"id"`
}

type ContentItem struct {
	QuestionType  string  `json:"question_type"`
	Question      []Title `json:"question"`
	Answer        string  `json:"answer"`
	WordId        string  `json:"word_id"`
	Word          string  `json:"word"`
	Definition    string  `json:"definition"`
	Pronunciation string  `json:"pronunciation"`
}

type ContentItems struct {
	ContentItems []ContentItem `json:"content_items"`
}
