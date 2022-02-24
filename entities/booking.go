package entities

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	RoomsID uint
	UserID  uint
	// Price          int    `gorm:"type:int(11)" json:"price"`
	DateTime int `gorm:"type:date" json:"date_time"`
	// TransactionsID uint
}
