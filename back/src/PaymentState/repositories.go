package paymentstate

import (
	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
)

func GetAllPaymentStatesRepository() ([]schema.PaymentState, error) {
	var paymentStates []schema.PaymentState
	db := config.DB
	err := db.Find(&paymentStates).Error
	return paymentStates, err
}
