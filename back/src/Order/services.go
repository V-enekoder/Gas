package order

import (
	"errors"

	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
	"gorm.io/gorm"
)

var (
	ErrCylinderNotFound = errors.New("one or more cylinder types not found")
)

const (
	// Suponiendo que el ID 1 es para el estado 'Pendiente' o 'Creado'
	DefaultOrderStateID = 1
)

func mapToOrderResponseDTO(o schema.Order) OrderResponseDTO {
	detailsDTO := make([]OrderDetailResponseDTO, len(o.OrderDetails))
	for i, d := range o.OrderDetails {
		detailsDTO[i] = OrderDetailResponseDTO{
			ID:       d.ID,
			Quantity: d.Quantity,
			Price:    d.Price,
			TypeCylinder: TypeCylinderInDetailResponseDTO{
				ID:   d.TypeCylinder.ID,
				Name: d.TypeCylinder.Name,
			},
		}
	}

	return OrderResponseDTO{
		ID:         o.ID,
		TotalPrice: o.TotalPrice,
		CreatedAt:  o.CreatedAt,
		User: UserInOrderResponseDTO{
			ID:    o.User.ID,
			Name:  o.User.Name,
			Email: o.User.Email,
		},
		OrderState: OrderStateInOrderResponseDTO{
			ID:   o.OrderState.ID,
			Name: o.OrderState.Name,
		},
		OrderDetails: detailsDTO,
	}
}

func CreateOrderService(dto OrderCreateDTO) (OrderResponseDTO, error) {
	var newOrder schema.Order
	var totalPrice float64 = 0

	// Usamos una transacción para la validación de precios y la creación.
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		for _, detailDTO := range dto.OrderDetails {
			cylinder, err := GetTypeCylinderForOrder(tx, detailDTO.TypeCylinderID)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return ErrCylinderNotFound
				}
				return err
			}

			detailPrice := cylinder.Price * float64(detailDTO.Quantity)
			totalPrice += detailPrice

			newOrder.OrderDetails = append(newOrder.OrderDetails, schema.OrderDetail{
				TypeCylinderID: detailDTO.TypeCylinderID,
				Quantity:       detailDTO.Quantity,
				Price:          detailPrice, // Guardar el precio total del detalle en ese momento
			})
		}
		return nil
	})

	if err != nil {
		return OrderResponseDTO{}, err
	}

	newOrder.UserID = dto.UserID
	newOrder.TotalPrice = totalPrice
	newOrder.StateOrderID = DefaultOrderStateID // Estado inicial por defecto

	if err := CreateOrderRepository(&newOrder); err != nil {
		return OrderResponseDTO{}, err
	}

	// Recuperamos la orden creada con todas las relaciones para la respuesta
	createdOrder, err := GetOrderByIDRepository(newOrder.ID)
	if err != nil {
		return OrderResponseDTO{}, err
	}

	return mapToOrderResponseDTO(createdOrder), nil
}

func GetAllOrdersService() ([]OrderResponseDTO, error) {
	orders, err := GetAllOrdersRepository()
	if err != nil {
		return nil, err
	}

	var responseDTOs []OrderResponseDTO
	for _, o := range orders {
		responseDTOs = append(responseDTOs, mapToOrderResponseDTO(o))
	}
	return responseDTOs, nil
}

func GetOrderByIDService(id uint) (OrderResponseDTO, error) {
	order, err := GetOrderByIDRepository(id)
	if err != nil {
		return OrderResponseDTO{}, err
	}
	return mapToOrderResponseDTO(order), nil
}

func GetOrdersByUserIDService(id uint) ([]OrderResponseDTO, error) {
	orders, err := GetOrdersByUserIDRepository(id)
	if err != nil {
		return []OrderResponseDTO{}, err
	}

	var ordersDTO []OrderResponseDTO

	for _, order := range orders {
		ordersDTO = append(ordersDTO, mapToOrderResponseDTO(order))
	}

	return ordersDTO, nil
}
