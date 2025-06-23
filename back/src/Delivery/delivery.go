package delivery

import "time"

type DeliveryCreateDTO struct {
	OrderID   uint `json:"order_id" binding:"required"`
	PaymentID uint `json:"payment_id" binding:"required"`
}

// --- DTOs de Respuesta ---

type TypeCylinderInDetailResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type DeliveryDetailResponseDTO struct {
	ID           uint                            `json:"id"`
	Quantity     int                             `json:"quantity"`
	TypeCylinder TypeCylinderInDetailResponseDTO `json:"type_cylinder"`
}

type OrderInDeliveryResponseDTO struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type PaymentInDeliveryResponseDTO struct {
	ID       uint      `json:"id"`
	Quantity float64   `json:"quantity"`
	Date     time.Time `json:"date"`
}

type DeliveryResponseDTO struct {
	ID              uint                         `json:"id"`
	TotalPrice      float64                      `json:"total_price"`
	Order           OrderInDeliveryResponseDTO   `json:"order"`
	Payment         PaymentInDeliveryResponseDTO `json:"payment"`
	DeliveryDetails []DeliveryDetailResponseDTO  `json:"delivery_details"`
}
