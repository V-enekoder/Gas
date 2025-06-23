package orderstate

import (
	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
)

func GetAllOrderStatesRepository() ([]schema.OrderState, error) {
	var orderStates []schema.OrderState
	db := config.DB
	err := db.Find(&orderStates).Error
	return orderStates, err
}
