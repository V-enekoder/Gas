package municipality

import (
	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
)

// GetAllMunicipalitiesRepository obtiene todos los municipios de la base de datos.
func GetAllMunicipalitiesRepository() ([]schema.Municipality, error) {
	var municipalities []schema.Municipality
	db := config.DB
	err := db.Find(&municipalities).Error
	return municipalities, err
}
