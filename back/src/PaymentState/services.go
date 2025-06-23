package paymentstate

import (
	"github.com/V-enekoder/GasManager/src/schema"
)

func mapToPaymentStateResponseDTO(p schema.PaymentState) PaymentStateResponseDTO {
	return PaymentStateResponseDTO{
		ID:   p.ID,
		Name: p.Name,
	}
}

func GetAllPaymentStatesService() ([]PaymentStateResponseDTO, error) {
	paymentStates, err := GetAllPaymentStatesRepository()
	if err != nil {
		return nil, err
	}

	var responseDTOs []PaymentStateResponseDTO
	for _, p := range paymentStates {
		responseDTOs = append(responseDTOs, mapToPaymentStateResponseDTO(p))
	}
	return responseDTOs, nil
}
