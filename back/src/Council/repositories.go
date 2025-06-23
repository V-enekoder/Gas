package council

import (
	"errors"

	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
	"gorm.io/gorm"
)

func CouncilExistsByUserIDRepository(userID uint) (bool, error) {
	var council schema.Council
	db := config.DB
	err := db.Where("user_id = ?", userID).First(&council).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CouncilExistsByLeaderDocumentRepository(document string, idToExclude uint) (bool, error) {
	var council schema.Council
	db := config.DB
	query := db.Where("leader_document = ?", document)
	if idToExclude > 0 {
		query = query.Where("id != ?", idToExclude)
	}
	err := query.First(&council).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CreateCouncilRepository(council *schema.Council) error {
	db := config.DB
	return db.Create(council).Error
}

func GetAllCouncilsRepository() ([]schema.Council, error) {
	var councils []schema.Council
	db := config.DB
	err := db.Preload("User").Find(&councils).Error
	return councils, err
}

func GetCouncilByIDRepository(id uint) (schema.Council, error) {
	var council schema.Council
	db := config.DB
	err := db.Preload("User").First(&council, id).Error
	return council, err
}

func UpdateCouncilRepository(id uint, data map[string]interface{}) error {
	db := config.DB
	result := db.Model(&schema.Council{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func DeleteCouncilRepository(id uint) error {
	db := config.DB
	result := db.Delete(&schema.Council{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
