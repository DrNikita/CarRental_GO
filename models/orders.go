package models

import "gorm.io/gorm"

type Orders struct{}

func MigrateOrders(db *gorm.DB) error {
	err := db.AutoMigrate(&Orders{})
	return err
}
