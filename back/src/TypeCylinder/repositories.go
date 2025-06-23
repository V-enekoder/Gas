package typecylinder

import (
	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
)

func GetAllTypeCylindersRepository() ([]schema.TypeCylinder, error) {
	var typeCylinders []schema.TypeCylinder
	db := config.DB
	err := db.Find(&typeCylinders).Error
	return typeCylinders, err
}
