package repository

import (
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestWordRepository_CreateWord(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Word{})
	repo := NewWordRepositoryImpl(mockDB)

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

	err := repo.CreateWord(word)
	assert.Nil(t, err)
}

func TestWordRepository_UpdateWord(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Word{})
	repo := NewWordRepositoryImpl(mockDB)

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
	result := repo.UpdateWord(word)
	assert.Nil(t, result)
	var words []models.Word
	mockDB.Find(&words)
	assert.Equal(t, 1, len(words))
	assert.Equal(t, "updated word", words[0].Word)
}

func TestWordRepository_DeleteWord(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Word{})
	repo := NewWordRepositoryImpl(mockDB)

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

	result := repo.DeleteWord("1")

	assert.Nil(t, result)
	var words []models.Word
	mockDB.Find(&words)
	assert.Equal(t, 0, len(words))
}

func TestWordRepository_GetWordByMemberIDANDUnitID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Word{})
	repo := NewWordRepositoryImpl(mockDB)

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

	words, result := repo.GetWordByUnitID("1")
	assert.Nil(t, result)
	assert.Equal(t, 1, len(words))
	assert.Equal(t, "word", words[0].Word)
}
