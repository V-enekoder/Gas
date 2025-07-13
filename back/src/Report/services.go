package report

import (
	"time"

	"github.com/V-enekoder/GasManager/src/schema"
)

func mapToReportResponseDTO(r schema.Report) ReportResponseDTO {
	return ReportResponseDTO{
		ID:          r.ID,
		UserID:      r.UserID,
		Description: r.Description,
		Date:        r.Date,
		Delivery: DeliveryInReportResponseDTO{
			ID:      r.Delivery.ID,
			OrderID: r.Delivery.OrderID,
		},
		ReportType: ReportTypeInResponseDTO{
			ID:   r.ReportType.ID,
			Name: r.ReportType.Name,
		},
		ReportState: ReportStateInResponseDTO{
			ID:   r.ReportState.ID,
			Name: r.ReportState.Name,
		},
	}
}

func CreateReportService(dto ReportCreateDTO) (ReportResponseDTO, error) {
	newReport := schema.Report{
		UserID:        dto.UserID,
		DeliveryID:    dto.DeliveryID,
		Description:   dto.Description,
		TypeID:        dto.TypeID,
		ReportStateID: dto.ReportStateID,
		Date:          time.Now(),
	}

	if err := CreateReportRepository(&newReport); err != nil {
		return ReportResponseDTO{}, err
	}

	createdReport, err := GetReportByIDRepository(newReport.ID)
	if err != nil {
		return ReportResponseDTO{}, err
	}

	return mapToReportResponseDTO(createdReport), nil
}

func GetAllReportsService() ([]ReportResponseDTO, error) {
	reports, err := GetAllReportsRepository()
	if err != nil {
		return nil, err
	}

	var responseDTOs []ReportResponseDTO
	for _, r := range reports {
		responseDTOs = append(responseDTOs, mapToReportResponseDTO(r))
	}
	return responseDTOs, nil
}

func GetReportByIDService(id uint) (ReportResponseDTO, error) {
	report, err := GetReportByIDRepository(id)
	if err != nil {
		return ReportResponseDTO{}, err
	}
	return mapToReportResponseDTO(report), nil
}

func GetReportByUserIDService(id uint) ([]ReportResponseDTO, error) {
	reports, err := GetReportByUserIDRepository(id)
	if err != nil {
		return []ReportResponseDTO{}, err
	}

	var reportsDTO []ReportResponseDTO
	for _, r := range reports {
		reportsDTO = append(reportsDTO, mapToReportResponseDTO(r))
	}

	return reportsDTO, nil
}
