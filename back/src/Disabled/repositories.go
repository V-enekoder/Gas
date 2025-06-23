package disabled

import (
	"errors"

	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
	"gorm.io/gorm"
)

func DisabledExistsByUserIDRepository(userID uint) (bool, error) {
	var disabled schema.Disabled
	db := config.DB
	err := db.Where("user_id = ?", userID).First(&disabled).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CreateDisabledRepository(disabled *schema.Disabled) error {
	db := config.DB
	return db.Create(disabled).Error
}

func GetAllDisabledRepository() ([]schema.Disabled, error) {
	var disabledList []schema.Disabled
	db := config.DB
	err := db.Preload("User").Find(&disabledList).Error
	return disabledList, err
}

func GetDisabledByIDRepository(id uint) (schema.Disabled, error) {
	var disabled schema.Disabled
	db := config.DB
	err := db.Preload("User").First(&disabled, id).Error
	return disabled, err
}

func UpdateDisabledRepository(id uint, data map[string]interface{}) error {
	db := config.DB
	result := db.Model(&schema.Disabled{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func DeleteDisabledRepository(id uint) error {
	db := config.DB
	result := db.Delete(&schema.Disabled{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
