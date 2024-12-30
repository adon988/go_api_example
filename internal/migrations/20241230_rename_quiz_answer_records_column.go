package migrations

import (
	"gorm.io/gorm"
)

func RenameCorrectAnswerCountToFailedAnswerCount(db *gorm.DB) error {
	err := db.Migrator().RenameColumn("quiz_answer_records", "correct_answer_count", "failed_answer_count")
	if err != nil {
		return err
	}
	return nil
}
