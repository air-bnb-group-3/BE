package entities

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	RoomsID        uint
	UserID         uint
	CheckIn        datatypes.Date //string `gorm:"type:datetime"`
	CheckOut       datatypes.Date //string `gorm:"type:datetime"`
	PaymentMethods string         `gorm:"type:enum('klikbca');default:'klikbca'"`
	Status         string         `gorm:"type:enum('payed','cancel','waiting');default:'waiting'"`
}
