package entities

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	RoomsID        uint
	UserID         uint
	Check_in       datatypes.Date
	Check_out      datatypes.Date
	PaymentMethods string `gorm:"type:enum('klikbca');default:'klikbca'"`
	Status         string `gorm:"type:enum('payed','cancel','waiting');default:'waiting'"`
}
