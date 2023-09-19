package repository

import "time"

type User struct {
	Id          uint       `gorm:"primary key;autoIncrement" json:"id"`
	RoleId      uint       `json:"role_id,omitempty" gorm:"foreignKey:RoleID;references:id"`
	FirstName   *string    `json:"first_name,omitempty"`
	SecondName  *string    `json:"second_name,omitempty"`
	Email       *string    `json:"email,omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
	Adress      *string    `json:"adress,omitempty"`
	BirthDate   *time.Time `json:"birth_date,omitempty"`
}
