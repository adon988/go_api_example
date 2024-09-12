package repository

import (
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUnitRepository_CreateUnit(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Unit{}, &models.UnitPermission{})
	repo := NewUnitRepository(mockDB)
	memberId := "1"

	unit := models.Unit{
		Id:        "1",
		Title:     "unit title",
		Order:     1,
		Publish:   1,
		CourseId:  "1",
		CreaterId: memberId,
	}

	err := repo.CreateUnit(unit)
	assert.Nil(t, err)
}

func TestUnitRepsitory_UpdateUnit(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Unit{}, &models.UnitPermission{})
	repo := NewUnitRepository(mockDB)
	memberId := "1"

	unit := models.Unit{
		Id:        "1",
		Title:     "unit title",
		Order:     1,
		Publish:   1,
		CourseId:  "1",
		CreaterId: memberId,
	}

	err := mockDB.Create(&unit)
	assert.Nil(t, err.Error)

	unit.Title = "updated unit title"
	result := repo.UpdateUnit(unit)
	assert.Nil(t, result)
	var units []models.Unit
	mockDB.Find(&units)
	assert.Equal(t, 1, len(units))
	assert.Equal(t, "updated unit title", units[0].Title)
}
func TestUnitRepsitory_DeleteUnit(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Unit{}, &models.UnitPermission{})
	repo := NewUnitRepository(mockDB)
	memberId := "1"

	unit := models.Unit{
		Id:        "1",
		Title:     "unit title",
		Order:     1,
		Publish:   1,
		CourseId:  "1",
		CreaterId: memberId,
	}

	err := mockDB.Create(&unit)
	assert.Nil(t, err.Error)

	result := repo.DeleteUnit("1")
	assert.Nil(t, result)
	var units []models.Unit
	mockDB.Find(&units)
	assert.Equal(t, 0, len(units))
}
func TestUnitRepsitory_GetUnitByMemberID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Unit{}, &models.UnitPermission{})
	repo := NewUnitRepository(mockDB)
	memberId := "1"

	unit := models.Unit{
		Id:        "1",
		Title:     "unit title",
		CourseId:  "1",
		Order:     1,
		Publish:   1,
		CreaterId: memberId,
	}

	mockDB.Create(&unit)

	unit_perm := models.UnitPermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	mockDB.Create(&unit_perm)

	units, err := repo.GetUnitByMemberID(memberId)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(units))
	assert.Equal(t, "unit title", units[0].Title)
}
