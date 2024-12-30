package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type QuizAnswerRecord struct {
	Id                 string           `gorm:"primaryKey;size:24"`
	QuizId             string           `gorm:"size:255;index;comment:'quiz id'"`
	MemberId           string           `gorm:"size:24;index;comment:'member id'"`
	AnswerQuestion     *json.RawMessage `gorm:"type:json;comment:'answer question'"`
	Status             int32            `gorm:"type:tinyint;comment:'status 1: unstart, 2: progress, 3: finish, 4: failed';default:1"`
	DueDate            *time.Time       `gorm:"type:date;comment:'due date'"`
	FailedAnswerCount  *int32           `gorm:"size:8;comment:'failed answer count';default:0"`
	TotalQuestionCount *int32           `gorm:"size:8;comment:'total answer count';default:0"`
	FailedLogs         *json.RawMessage `gorm:"type:json;comment:'failed logs'"`
	Scope              *int32           `gorm:"size:8;comment:'when finish quiz, caculate the persional scope';"`
	CreatedAt          time.Time        `gorm:"autoCreateTime"`
	UpdatedAt          time.Time        `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt   `gorm:"index"`
}
