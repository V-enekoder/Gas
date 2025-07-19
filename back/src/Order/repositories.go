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

func CreateOrderRepository(order *schema.Order) error {
	db := config.DB
	return db.Create(order).Error
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

func GetOrdersByUserIDRepository(userID uint) ([]schema.Order, error) {
	// 1. Declara una slice de órdenes, no una sola.
	var orders []schema.Order
	db := config.DB

	// 2. Usa Where() para filtrar por user_id y Find() para obtener múltiples registros.
	err := db.
		Preload("OrderState").                // Precarga el estado de cada orden
		Preload("OrderDetails").              // Precarga los detalles de cada orden
		Preload("OrderDetails.TypeCylinder"). // Precarga el tipo de cilindro DENTRO de cada detalle
		Where("user_id = ?", userID).         // La condición de filtrado clave
		Order("created_at DESC").             // Opcional: ordena las órdenes de la más nueva a la más antigua
		Find(&orders).Error                   // Busca todos los registros que coincidan y los carga en la slice

	return orders, err
}

func UpdateOrderStateRepository(orderID uint, newStateID uint) error {
	db := config.DB

	result := db.Model(&schema.Order{}).Where("id = ?", orderID).Update("state_order_id", newStateID)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
