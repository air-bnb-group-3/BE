package entities

import (
	"gorm.io/gorm"
)

type Midtrans struct {
	gorm.Model
	Booking []Booking `gorm:"ForeignKey:MidtransID"`
}
