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

func TestWordRepository_GetWordByMemberIDAndUnitID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Word{}, &models.UnitPermission{})
	repo := NewWordRepositoryImpl(mockDB)
	member_id := "1"

	unit_perm := models.UnitPermission{
		MemberId: "1",
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

	words, result := repo.GetWordByMemberIDAndUnitID(member_id, "1")
	assert.Nil(t, result)
	assert.Equal(t, 1, len(words))
	assert.Equal(t, "word", words[0].Word)
}

func TestWordRepository_CheckWordPermissionByMemberIDAndWordID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Word{}, &models.UnitPermission{})
	repo := NewWordRepositoryImpl(mockDB)
	member_id := "1"

	unit_perm := models.UnitPermission{
		MemberId: "1",
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

	wordData, result := repo.CheckWordPermissionByMemberIDAndWordID(member_id, word.Id)
	assert.Nil(t, result)
	assert.Equal(t, word.Word, wordData.Word)
}

func TestWordRepository_GetWordByMemberIDAndCourseID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Word{}, &models.Unit{}, &models.UnitPermission{}, &models.CoursePermission{})
	repo := NewWordRepositoryImpl(mockDB)
	member_id := "1"
	unit := models.Unit{
		Id:        "1",
		Title:     "unit title",
		Order:     1,
		Publish:   1,
		CourseId:  "1",
		CreaterId: member_id,
	}

	unit_perm := models.UnitPermission{
		MemberId: member_id,
		EntityId: "1",
		Role:     1,
	}

	course_perm := models.CoursePermission{
		MemberId: member_id,
		EntityId: "1",
		Role:     1,
	}
	mockDB.Create(&course_perm)
	mockDB.Create(&unit_perm)
	mockDB.Create(&unit)
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

	words, result := repo.GetWordByMemberIDAndCourseID(member_id, "1")
	assert.Nil(t, result)
	assert.Equal(t, 1, len(words))
	assert.Equal(t, "word", words[0].Word)
}

func TestWordRepository_Fail_GetWordByMemberIDAndCourseID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Word{}, &models.Unit{}, &models.UnitPermission{}, &models.CoursePermission{})
	repo := NewWordRepositoryImpl(mockDB)
	member_id := "1"
	member_id2 := "2"
	mockDB.Create(&models.CoursePermission{
		MemberId: member_id,
		EntityId: "1",
		Role:     1,
	})
	mockDB.Create(&models.UnitPermission{
		MemberId: member_id2,
		EntityId: "1",
		Role:     1,
	})
	mockDB.Create(&models.Unit{
		Id:        "1",
		Title:     "unit title",
		Order:     1,
		Publish:   1,
		CourseId:  "1",
		CreaterId: member_id,
	})
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

	words, result := repo.GetWordByMemberIDAndCourseID(member_id, "1")
	assert.Nil(t, result)
	assert.Equal(t, 0, len(words))
}
