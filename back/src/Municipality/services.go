package municipality

import (
	"github.com/V-enekoder/GasManager/src/schema"
)

func mapToMunicipalityResponseDTO(m schema.Municipality) MunicipalityResponseDTO {
	return MunicipalityResponseDTO{
		ID:   m.ID,
		Name: m.Name,
	}
}

// GetAllMunicipalitiesService recupera todos los municipios.
func GetAllMunicipalitiesService() ([]MunicipalityResponseDTO, error) {
	municipalities, err := GetAllMunicipalitiesRepository()
	if err != nil {
		return nil, err
	}

	var responseDTOs []MunicipalityResponseDTO
	for _, m := range municipalities {
		responseDTOs = append(responseDTOs, mapToMunicipalityResponseDTO(m))
	}
	return responseDTOs, nil
}
