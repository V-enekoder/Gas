package orderstate

import (
	"github.com/V-enekoder/GasManager/src/schema"
)

func mapToOrderStateResponseDTO(o schema.OrderState) OrderStateResponseDTO {
	return OrderStateResponseDTO{
		ID:   o.ID,
		Name: o.Name,
	}
}

func GetAllOrderStatesService() ([]OrderStateResponseDTO, error) {
	orderStates, err := GetAllOrderStatesRepository()
	if err != nil {
		return nil, err
	}

	var responseDTOs []OrderStateResponseDTO
	for _, o := range orderStates {
		responseDTOs = append(responseDTOs, mapToOrderStateResponseDTO(o))
	}
	return responseDTOs, nil
}
