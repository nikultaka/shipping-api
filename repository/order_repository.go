package repository

import (
	"shipping-api/db"
	"shipping-api/models"
)

func OrderExists(orderID string) (bool, error) {
	var count int64
	err := db.DB.Model(&models.Order{}).Where("order_id = ?", orderID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateOrder(order *models.Order) error {
	return db.DB.Create(order).Error
}
