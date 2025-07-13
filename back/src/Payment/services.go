package payment

import (
	"github.com/V-enekoder/GasManager/src/schema"
)

func mapToPaymentResponseDTO(p schema.Payment) PaymentResponseDTO {
	var deliveryID *uint
	if p.Delivery != nil {
		deliveryID = &p.Delivery.ID
	}

	return PaymentResponseDTO{
		ID:       p.ID,
		UserID:   p.UserID,
		OrderID:  p.OrderID,
		Quantity: p.Quantity,
		PaymentState: PaymentStateInPaymentResponseDTO{
			ID:   p.PaymentState.ID,
			Name: p.PaymentState.Name,
		},
		DeliveryID: deliveryID,
	}
}

func CreatePaymentService(dto PaymentCreateDTO) (PaymentResponseDTO, error) {
	newPayment := schema.Payment{
		UserID:   dto.UserID,
		OrderID:  dto.OrderID,
		Quantity: dto.Quantity,
		StateID:  1,
	}

	if err := CreatePaymentRepository(&newPayment); err != nil {
		return PaymentResponseDTO{}, err
	}

	createdPayment, err := GetPaymentByIDRepository(newPayment.ID)
	if err != nil {
		return PaymentResponseDTO{}, err
	}

	return mapToPaymentResponseDTO(createdPayment), nil
}

func GetAllPaymentsService() ([]PaymentResponseDTO, error) {
	payments, err := GetAllPaymentsRepository()
	if err != nil {
		return nil, err
	}

	var responseDTOs []PaymentResponseDTO
	for _, p := range payments {
		responseDTOs = append(responseDTOs, mapToPaymentResponseDTO(p))
	}
	return responseDTOs, nil
}

func GetPaymentByIDService(id uint) (PaymentResponseDTO, error) {
	payment, err := GetPaymentByIDRepository(id)
	if err != nil {
		return PaymentResponseDTO{}, err
	}
	return mapToPaymentResponseDTO(payment), nil
}
func GetPaymentByUserIDService(id uint) ([]PaymentResponseDTO, error) {
	payments, err := GetPaymentByUserIDRepository(id)
	if err != nil {
		return []PaymentResponseDTO{}, err
	}

	var paymentsDTO []PaymentResponseDTO

	for _, payment := range payments {
		paymentsDTO = append(paymentsDTO, mapToPaymentResponseDTO(payment))
	}

	return paymentsDTO, nil
}
