package order

import (
	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
	"gorm.io/gorm"
)

// GetTypeCylinderForOrder es una función auxiliar para obtener información del cilindro durante la creación de la orden.
func GetTypeCylinderForOrder(tx *gorm.DB, id uint) (schema.TypeCylinder, error) {
	var cylinder schema.TypeCylinder
	err := tx.First(&cylinder, id).Error
	return cylinder, err
}

// CreateOrderRepository crea una orden y sus detalles dentro de una transacción.
func CreateOrderRepository(order *schema.Order) error {
	db := config.DB
	return db.Transaction(func(tx *gorm.DB) error {
		// 1. Crear la Orden principal
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		// 2. Crear los Detalles de la Orden, asignando el OrderID recién creado
		for i := range order.OrderDetails {
			order.OrderDetails[i].OrderID = order.ID
			if err := tx.Create(&order.OrderDetails[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// GetAllOrdersRepository obtiene todas las órdenes con sus relaciones precargadas.
func GetAllOrdersRepository() ([]schema.Order, error) {
	var orders []schema.Order
	db := config.DB
	err := db.Preload("User").
		Preload("OrderState").
		Preload("OrderDetails.TypeCylinder").
		Order("created_at desc").
		Find(&orders).Error
	return orders, err
}

// GetOrderByIDRepository obtiene una orden por su ID con sus relaciones precargadas.
func GetOrderByIDRepository(id uint) (schema.Order, error) {
	var order schema.Order
	db := config.DB
	err := db.Preload("User").
		Preload("OrderState").
		Preload("OrderDetails.TypeCylinder").
		First(&order, id).Error
	return order, err
}
