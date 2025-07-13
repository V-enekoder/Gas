package payment

import "time"

type PaymentCreateDTO struct {
	UserID   uint    `json:"user_id" binding:"required"`
	OrderID  uint    `json:"order_id" binding:"required"`
	Quantity float64 `json:"quantity" binding:"required"`
}

type PaymentStateInPaymentResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PaymentResponseDTO struct {
	ID           uint                             `json:"id"`
	UserID       uint                             `json:"user_id"`
	OrderID      uint                             `json:"order_id"`
	Quantity     float64                          `json:"quantity"`
	Date         time.Time                        `json:"date"`
	PaymentState PaymentStateInPaymentResponseDTO `json:"payment_state"`
	DeliveryID   *uint                            `json:"delivery_id,omitempty"`
}
