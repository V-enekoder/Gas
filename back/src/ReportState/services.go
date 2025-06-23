package reportstate

import (
	"github.com/V-enekoder/GasManager/src/schema"
)

func mapToReportStateResponseDTO(r schema.ReportState) ReportStateResponseDTO {
	return ReportStateResponseDTO{
		ID:   r.ID,
		Name: r.Name,
	}
}

func GetAllReportStatesService() ([]ReportStateResponseDTO, error) {
	reportStates, err := GetAllReportStatesRepository()
	if err != nil {
		return nil, err
	}

	var responseDTOs []ReportStateResponseDTO
	for _, r := range reportStates {
		responseDTOs = append(responseDTOs, mapToReportStateResponseDTO(r))
	}
	return responseDTOs, nil
}
