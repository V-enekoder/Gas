package reportstate

import (
	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
)

func GetAllReportStatesRepository() ([]schema.ReportState, error) {
	var reportStates []schema.ReportState
	db := config.DB
	err := db.Find(&reportStates).Error
	return reportStates, err
}
