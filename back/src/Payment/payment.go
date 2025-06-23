package payment

import "time"

type PaymentCreateDTO struct {
	Quantity float64 `json:"quantity" binding:"required"`
	StateID  uint    `json:"state_id" binding:"required"`
}

type PaymentStateInPaymentResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PaymentResponseDTO struct {
	ID           uint                             `json:"id"`
	Quantity     float64                          `json:"quantity"`
	Date         time.Time                        `json:"date"`
	PaymentState PaymentStateInPaymentResponseDTO `json:"payment_state"`
	DeliveryID   *uint                            `json:"delivery_id,omitempty"`
}
