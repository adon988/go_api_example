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

func TestUnitRepository_GetUnitsByCourseID(t *testing.T) {
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

	units, result := repo.GetUnitsByCourseID("1")
	assert.Nil(t, result)
	assert.Equal(t, 1, len(units))
	assert.Equal(t, "unit title", units[0].Title)
}

func TestUnitRepository_GetUnitByMemberIDAndUnitID(t *testing.T) {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	mockDB.AutoMigrate(&models.Unit{}, &models.UnitPermission{})
	repo := NewUnitRepository(mockDB)
	memberId := "1"
	unAuthMemberId := "2"
	entityId := "1"
	public := int32(1)
	unPublish := int32(0)
	unit := models.Unit{
		Id:        "1",
		Title:     "unit title",
		Order:     1,
		Publish:   public,
		CourseId:  "1",
		CreaterId: memberId,
	}

	unit_perm := models.UnitPermission{
		MemberId: "1",
		EntityId: "1",
		Role:     1,
	}
	mockDB.Create(&unit_perm)
	mockDB.Create(&unit)

	//check public scenario by auth member
	unitData, err := repo.GetUnitByMemberIDAndUnitID(memberId, entityId)
	assert.Nil(t, err)
	assert.Equal(t, unit.Id, unitData.Id)
	assert.Equal(t, unit.Title, unitData.Title)

	//check public scenario by un-auth member
	unitData, err = repo.GetUnitByMemberIDAndUnitID(unAuthMemberId, entityId)
	assert.Nil(t, err)
	assert.Equal(t, unit.Id, unitData.Id)
	assert.Equal(t, unit.Title, unitData.Title)

	// change public status to un-publish
	unit.Publish = unPublish
	mockDB.Save(&unit)

	// check private status by auth member
	unitData, err = repo.GetUnitByMemberIDAndUnitID(memberId, entityId)
	assert.Nil(t, err)
	assert.Equal(t, unitData.Id, unit.Id)
	assert.Equal(t, unitData.Title, unit.Title)

	//check non-public scenario by un-auth member
	unitData, err = repo.GetUnitByMemberIDAndUnitID(unAuthMemberId, entityId)
	assert.Error(t, err)
	assert.Equal(t, unitData, models.Unit{})

}
