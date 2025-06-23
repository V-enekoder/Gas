package reporttype

import (
	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
)

func GetAllReportTypesRepository() ([]schema.ReportType, error) {
	var reportTypes []schema.ReportType
	db := config.DB
	err := db.Find(&reportTypes).Error
	return reportTypes, err
}
