package models

import "gorm.io/gorm"

type Damages struct{}

func MigrateDamages(db *gorm.DB) error {
	err := db.AutoMigrate(&Damages{})
	return err
}
