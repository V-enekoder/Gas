package delivery

import (
	"errors"

	order "github.com/V-enekoder/GasManager/src/Order"
	"github.com/V-enekoder/GasManager/src/schema"
	"gorm.io/gorm"
)

var (
	ErrOrderNotFound         = errors.New("the source order was not found")
	ErrDeliveryAlreadyExists = errors.New("a delivery for this order already exists")
)

const (
	// Suponiendo que el ID 1 en PaymentState es 'Pagado' o 'Completo'
	DefaultPaymentStateID = 1
)

func mapToDeliveryResponseDTO(d schema.Delivery) DeliveryResponseDTO {
	detailsDTO := make([]DeliveryDetailResponseDTO, len(d.DeliveryDetails))
	for i, detail := range d.DeliveryDetails {
		detailsDTO[i] = DeliveryDetailResponseDTO{
			ID:       detail.ID,
			Quantity: detail.Quantity,
			TypeCylinder: TypeCylinderInDetailResponseDTO{
				ID:   detail.TypeCylinder.ID,
				Name: detail.TypeCylinder.Name,
			},
		}
	}

	return DeliveryResponseDTO{
		ID:         d.ID,
		TotalPrice: d.TotalPrice,
		Order: OrderInDeliveryResponseDTO{
			ID:        d.Order.ID,
			CreatedAt: d.Order.CreatedAt,
		},
		Payment: PaymentInDeliveryResponseDTO{
			ID:       d.Payment.ID,
			Quantity: d.Payment.Quantity,
		},
		DeliveryDetails: detailsDTO,
	}
}

func CreateDeliveryService(dto DeliveryCreateDTO) (DeliveryResponseDTO, error) {
	exists, err := DeliveryExistsByOrderIDRepository(dto.OrderID)
	if err != nil {
		return DeliveryResponseDTO{}, err
	}
	if exists {
		return DeliveryResponseDTO{}, ErrDeliveryAlreadyExists
	}

	sourceOrder, err := GetSourceOrderByIDRepository(dto.OrderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return DeliveryResponseDTO{}, ErrOrderNotFound
		}
		return DeliveryResponseDTO{}, err
	}

	// Preparar el nuevo Delivery
	newDelivery := schema.Delivery{
		OrderID:    sourceOrder.ID,
		PaymentID:  dto.PaymentID,
		TotalPrice: sourceOrder.TotalPrice,
	}

	// Preparar los detalles del Delivery basados en los detalles de la orden
	for _, orderDetail := range sourceOrder.OrderDetails {
		newDelivery.DeliveryDetails = append(newDelivery.DeliveryDetails, schema.DeliveryDetail{
			TypeCylinderID: orderDetail.TypeCylinderID,
			Quantity:       orderDetail.Quantity,
		})
	}

	// Crear todo en una transacción
	if err := CreateDeliveryRepository(&newDelivery); err != nil {
		return DeliveryResponseDTO{}, err
	}

	// Recuperar el delivery completo para la respuesta
	createdDelivery, err := GetDeliveryByIDRepository(newDelivery.ID)
	if err != nil {
		return DeliveryResponseDTO{}, err
	}

	_, err = order.UpdateOrderStateService(sourceOrder.ID)

	return mapToDeliveryResponseDTO(createdDelivery), nil
}

func GetAllDeliveriesService() ([]DeliveryResponseDTO, error) {
	deliveries, err := GetAllDeliveriesRepository()
	if err != nil {
		return nil, err
	}

	var responseDTOs []DeliveryResponseDTO
	for _, d := range deliveries {
		responseDTOs = append(responseDTOs, mapToDeliveryResponseDTO(d))
	}
	return responseDTOs, nil
}

func GetDeliveryByIDService(id uint) (DeliveryResponseDTO, error) {
	delivery, err := GetDeliveryByIDRepository(id)
	if err != nil {
		return DeliveryResponseDTO{}, err
	}
	return mapToDeliveryResponseDTO(delivery), nil
}

func GetDeliveriesByUserIDService(id uint) ([]DeliveryResponseDTO, error) {
	deliveries, err := GetDeliveriesByUserIDRepository(id)
	if err != nil {
		return []DeliveryResponseDTO{}, err
	}

	var deliveriesDTO []DeliveryResponseDTO
	for _, d := range deliveries {
		deliveriesDTO = append(deliveriesDTO, mapToDeliveryResponseDTO(d))
	}

	return deliveriesDTO, nil
}
