package models

import (
	"time"

	"gorm.io/gorm"
)

type Cars struct {
	Id         uint       `gorm:"primary key;autoIncrement" json:"id"`
	Govnum     *string    `json:"govnum"`
	Brand      *string    `json:"brand"`
	IssueDate  *time.Time `json:"issue_date"`
	CarCost    uint       `json:"car_cost"`
	RentalCost uint       `json:"rental_cost"`
}

func MigrateCars(db *gorm.DB) error {
	err := db.AutoMigrate(&Cars{})
	return err
}
