package order

import "time"

// OrderDetailCreateDTO es parte de la solicitud de creación de una orden.
type OrderDetailCreateDTO struct {
	TypeCylinderID uint `json:"type_cylinder_id" binding:"required"`
	Quantity       int  `json:"quantity" binding:"required,gt=0"`
}

// OrderCreateDTO define la estructura para crear una nueva orden.
type OrderCreateDTO struct {
	UserID       uint                   `json:"user_id" binding:"required"`
	OrderDetails []OrderDetailCreateDTO `json:"order_details" binding:"required,min=1"`
}

// --- DTOs de Respuesta ---

type UserInOrderResponseDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type OrderStateInOrderResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type TypeCylinderInDetailResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type OrderDetailResponseDTO struct {
	ID           uint                            `json:"id"`
	Quantity     int                             `json:"quantity"`
	Price        float64                         `json:"price"`
	TypeCylinder TypeCylinderInDetailResponseDTO `json:"type_cylinder"`
}

type OrderResponseDTO struct {
	ID           uint                         `json:"id"`
	TotalPrice   float64                      `json:"total_price"`
	CreatedAt    time.Time                    `json:"created_at"`
	User         UserInOrderResponseDTO       `json:"user"`
	OrderState   OrderStateInOrderResponseDTO `json:"order_state"`
	OrderDetails []OrderDetailResponseDTO     `json:"order_details"`
}
