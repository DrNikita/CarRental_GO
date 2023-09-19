package models

import "gorm.io/gorm"

type Users struct{}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err
}
