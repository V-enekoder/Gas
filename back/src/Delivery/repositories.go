package delivery

import (
	"errors"

	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
	"gorm.io/gorm"
)

func DeliveryExistsByOrderIDRepository(orderID uint) (bool, error) {
	var delivery schema.Delivery
	db := config.DB
	err := db.Where("order_id = ?", orderID).First(&delivery).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func GetSourceOrderByIDRepository(id uint) (schema.Order, error) {
	var order schema.Order
	db := config.DB
	err := db.Preload("OrderDetails").First(&order, id).Error
	return order, err
}

func CreateDeliveryRepository(delivery *schema.Delivery) error {
	db := config.DB
	return db.Transaction(func(tx *gorm.DB) error {
		// 1. Crear el Pago asociado.
		/*if err := tx.Create(payment).Error; err != nil {
			return err
		}

		// 2. Asignar el ID del pago al delivery y crear el delivery.
		delivery.PaymentID = payment.ID*/
		if err := tx.Create(delivery).Error; err != nil {
			return err
		}

		// 3. Crear los Detalles del Delivery, asignando el DeliveryID recién creado.
		/*for i := range delivery.DeliveryDetails {
			delivery.DeliveryDetails[i].DeliveryID = delivery.ID
			if err := tx.Create(&delivery.DeliveryDetails[i]).Error; err != nil {
				return err
			}
		}*/

		return nil
	})
}

func GetAllDeliveriesRepository() ([]schema.Delivery, error) {
	var deliveries []schema.Delivery
	db := config.DB
	err := db.Preload("Order").
		Preload("Payment").
		Preload("DeliveryDetails.TypeCylinder").
		Find(&deliveries).Error
	return deliveries, err
}

func GetDeliveryByIDRepository(id uint) (schema.Delivery, error) {
	var delivery schema.Delivery
	db := config.DB
	err := db.Preload("Order").
		Preload("Payment").
		Preload("DeliveryDetails.TypeCylinder").
		First(&delivery, id).Error
	return delivery, err
}

func GetDeliveriesByUserIDRepository(userID uint) ([]schema.Delivery, error) {
	var deliveries []schema.Delivery
	db := config.DB

	err := db.
		Joins("JOIN orders ON orders.id = deliveries.order_id").
		Where("orders.user_id = ?", userID).
		Preload("Order").
		Preload("Payment").
		Preload("DeliveryDetails.TypeCylinder").
		Find(&deliveries).Error

	return deliveries, err
}

