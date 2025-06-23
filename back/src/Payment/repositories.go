package payment

import (
	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
)

func CreatePaymentRepository(payment *schema.Payment) error {
	db := config.DB
	return db.Create(payment).Error
}

func GetAllPaymentsRepository() ([]schema.Payment, error) {
	var payments []schema.Payment
	db := config.DB
	err := db.Preload("PaymentState").Preload("Delivery").Find(&payments).Error
	return payments, err
}

func GetPaymentByIDRepository(id uint) (schema.Payment, error) {
	var payment schema.Payment
	db := config.DB
	err := db.Preload("PaymentState").Preload("Delivery").First(&payment, id).Error
	return payment, err
}
