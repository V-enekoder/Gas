package user

import (
	"errors"

	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
	"gorm.io/gorm"
)

func UserExistsByEmailRepository(email string) (bool, error) {
	var user schema.User
	db := config.DB
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CreateUserRepository(user *schema.User) error {
	db := config.DB
	return db.Create(user).Error
}

func GetAllUsersRepository() ([]schema.User, error) {
	var users []schema.User
	db := config.DB
	err := db.Preload("Municipality").Find(&users).Error
	return users, err
}

func GetUserByIDRepository(id uint) (schema.User, error) {
	var user schema.User
	db := config.DB
	err := db.Preload("Municipality").First(&user, id).Error
	return user, err
}

func UpdateUserRepository(id uint, data map[string]interface{}) error {
	db := config.DB
	result := db.Model(&schema.User{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func DeleteUserRepository(id uint) error {
	db := config.DB

	// Iniciar una transacción para garantizar la atomicidad de las comprobaciones y la eliminación.
	return db.Transaction(func(tx *gorm.DB) error {
		var orderCount int64
		tx.Model(&schema.Order{}).Where("user_id = ?", id).Count(&orderCount)
		if orderCount > 0 {
			return ErrUserHasRelations
		}

		var commerceCount int64
		tx.Model(&schema.Commerce{}).Where("user_id = ?", id).Count(&commerceCount)
		if commerceCount > 0 {
			return ErrUserHasRelations
		}

		var DisabledCount int64
		tx.Model(&schema.Disabled{}).Where("user_id = ?", id).Count(&DisabledCount)
		if commerceCount > 0 {
			return ErrUserHasRelations
		}

		var CouncilCount int64
		tx.Model(&schema.Council{}).Where("user_id = ?", id).Count(&CouncilCount)
		if commerceCount > 0 {
			return ErrUserHasRelations
		}

		result := tx.Delete(&schema.User{}, id)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
		return nil
	})
}
