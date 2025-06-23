package typecylinder

import (
	"github.com/V-enekoder/GasManager/src/schema"
)

func mapToTypeCylinderResponseDTO(t schema.TypeCylinder) TypeCylinderResponseDTO {
	return TypeCylinderResponseDTO{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Price:       t.Price,
		Disponible:  t.Disponible,
	}
}

func GetAllTypeCylindersService() ([]TypeCylinderResponseDTO, error) {
	typeCylinders, err := GetAllTypeCylindersRepository()
	if err != nil {
		return nil, err
	}

	var responseDTOs []TypeCylinderResponseDTO
	for _, t := range typeCylinders {
		responseDTOs = append(responseDTOs, mapToTypeCylinderResponseDTO(t))
	}
	return responseDTOs, nil
}
