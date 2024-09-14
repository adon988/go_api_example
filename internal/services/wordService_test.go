package services

import (
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestWordService_CreateWrod(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Word{})
	service := NewWordService(mockDB)

	word := models.Word{
		Id:            "1",
		UnitId:        "1",
		Word:          "word",
		Definition:    "definition",
		Image:         "image",
		Pronunciation: "pronunciation",
		Description:   "description",
		Comment:       "comment",
		Order:         1,
	}

	err := service.CreateWord(word)
	assert.Nil(t, err)
}

func TestWordService_UpdateWord(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Word{})
	service := NewWordService(mockDB)

	word := models.Word{
		Id:            "1",
		UnitId:        "1",
		Word:          "word",
		Definition:    "definition",
		Image:         "image",
		Pronunciation: "pronunciation",
		Description:   "description",
		Comment:       "comment",
		Order:         1,
	}

	err := mockDB.Create(&word)
	assert.Nil(t, err.Error)

	word.Word = "updated word"
	result := service.UpdateWord(word)
	assert.Nil(t, result)
	var words []models.Word
	mockDB.Find(&words)
	assert.Equal(t, 1, len(words))
	assert.Equal(t, "updated word", words[0].Word)
}

func TestWordService_DeleteWord(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Word{})
	service := NewWordService(mockDB)
	word := models.Word{
		Id:            "1",
		UnitId:        "1",
		Word:          "word",
		Definition:    "definition",
		Image:         "image",
		Pronunciation: "pronunciation",
		Description:   "description",
		Comment:       "comment",
		Order:         1,
	}

	err := mockDB.Create(&word)
	assert.Nil(t, err.Error)

	result := service.DeleteWord("1")
	assert.Nil(t, result)
	var words []models.Word
	mockDB.Find(&words)
	assert.Equal(t, 0, len(words))
}

func TestWordService_GetWordByMemberIDAndUnitID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Word{}, &models.UnitPermission{})
	service := NewWordService(mockDB)
	member_id := "1"

	unit_perm := models.UnitPermission{
		MemberId: member_id,
		EntityId: "1",
		Role:     1,
	}
	mockDB.Create(&unit_perm)
	word := models.Word{
		Id:            "1",
		UnitId:        "1",
		Word:          "word",
		Definition:    "definition",
		Image:         "image",
		Pronunciation: "pronunciation",
		Description:   "description",
		Comment:       "comment",
		Order:         1,
	}

	err := mockDB.Create(&word)
	assert.Nil(t, err.Error)

	words, result := service.GetWordByMemberIDAndUnitID(member_id, word.UnitId)
	assert.Nil(t, result)
	assert.Equal(t, 1, len(words))
}

func TestWordService_CheckWordPermissionByMemberIDAndWordID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Word{}, &models.UnitPermission{})
	service := NewWordService(mockDB)
	member_id := "1"

	unit_perm := models.UnitPermission{
		MemberId: member_id,
		EntityId: "1",
		Role:     1,
	}
	mockDB.Create(&unit_perm)
	word := models.Word{
		Id:            "1",
		UnitId:        "1",
		Word:          "word",
		Definition:    "definition",
		Image:         "image",
		Pronunciation: "pronunciation",
		Description:   "description",
		Comment:       "comment",
		Order:         1,
	}

	err := mockDB.Create(&word)
	assert.Nil(t, err.Error)

	words, result := service.CheckWordPermissionByMemberIDAndWordID(member_id, word.Id)
	assert.Nil(t, result)
	assert.Equal(t, "word", words.Word)
}
