package entities

import (
	"gorm.io/gorm"
)

type PaymentMethods struct {
	gorm.Model
	NameMethods string    `gorm:"type:varchar(100)" json:"name_methods"`
	Booking     []Booking `gorm:"ForeignKey:PaymentMethodsID"`
}
