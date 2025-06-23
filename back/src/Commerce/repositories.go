package commerce

import (
	"errors"

	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
	"gorm.io/gorm"
)

func CommerceExistsByUserIDRepository(userID uint) (bool, error) {
	var commerce schema.Commerce
	db := config.DB
	err := db.Where("user_id = ?", userID).First(&commerce).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CommerceExistsByRifRepository(rif string, idToExclude uint) (bool, error) {
	var commerce schema.Commerce
	db := config.DB
	query := db.Where("rif = ?", rif)
	if idToExclude > 0 {
		query = query.Where("id != ?", idToExclude)
	}
	err := query.First(&commerce).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CommerceExistsByBossDocumentRepository(document string, idToExclude uint) (bool, error) {
	var commerce schema.Commerce
	db := config.DB
	query := db.Where("boss_document = ?", document)
	if idToExclude > 0 {
		query = query.Where("id != ?", idToExclude)
	}
	err := query.First(&commerce).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CreateCommerceRepository(commerce *schema.Commerce) error {
	db := config.DB
	return db.Create(commerce).Error
}

func GetAllCommercesRepository() ([]schema.Commerce, error) {
	var commerces []schema.Commerce
	db := config.DB
	err := db.Preload("User").Find(&commerces).Error
	return commerces, err
}

func GetCommerceByIDRepository(id uint) (schema.Commerce, error) {
	var commerce schema.Commerce
	db := config.DB
	err := db.Preload("User").First(&commerce, id).Error
	return commerce, err
}

func UpdateCommerceRepository(id uint, data map[string]interface{}) error {
	db := config.DB
	result := db.Model(&schema.Commerce{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func DeleteCommerceRepository(id uint) error {
	db := config.DB
	result := db.Delete(&schema.Commerce{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
