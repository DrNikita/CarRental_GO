package models

import "gorm.io/gorm"

type Roles struct{}

func MigrateRoles(db *gorm.DB) error {
	err := db.AutoMigrate(&Roles{})
	return err
}
