package entities

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	RoomsID          uint
	UserID           uint
	MidtransID       uint
	PaymentMethodsID uint
	Price            int    `gorm:"type:int(11)" json:"price"`
	Phone            string `gorm:"type:varchar(13); not null" json:"phone"`
	DateTime         string `gorm:"type:date" json:"date_time"`
	Status           string `gorm:"type:enum('Payed','Canceled')" json:"status"`
}
