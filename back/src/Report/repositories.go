package report

import (
	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
)

func CreateReportRepository(report *schema.Report) error {
	db := config.DB
	return db.Create(report).Error
}

func GetAllReportsRepository() ([]schema.Report, error) {
	var reports []schema.Report
	db := config.DB
	err := db.Preload("Delivery").
		Preload("ReportType").
		Preload("ReportState").
		Order("date desc").
		Find(&reports).Error
	return reports, err
}

func GetReportByIDRepository(id uint) (schema.Report, error) {
	var report schema.Report
	db := config.DB
	err := db.Preload("Delivery").
		Preload("ReportType").
		Preload("ReportState").
		First(&report, id).Error
	return report, err
}
