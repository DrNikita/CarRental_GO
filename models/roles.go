package models

import "gorm.io/gorm"

type Roles struct {
	Id       uint    `gorm:"primary key;autoIncrement" json:"id"`
	RoleName *string `json:"role_name,omitempty"`
}

func MigrateRoles(db *gorm.DB) error {
	err := db.AutoMigrate(&Roles{})
	return err
}
