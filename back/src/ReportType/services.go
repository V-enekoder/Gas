package reporttype

import (
	"github.com/V-enekoder/GasManager/src/schema"
)

func mapToReportTypeResponseDTO(r schema.ReportType) ReportTypeResponseDTO {
	return ReportTypeResponseDTO{
		ID:   r.ID,
		Name: r.Name,
	}
}

func GetAllReportTypesService() ([]ReportTypeResponseDTO, error) {
	reportTypes, err := GetAllReportTypesRepository()
	if err != nil {
		return nil, err
	}

	var responseDTOs []ReportTypeResponseDTO
	for _, r := range reportTypes {
		responseDTOs = append(responseDTOs, mapToReportTypeResponseDTO(r))
	}
	return responseDTOs, nil
}
