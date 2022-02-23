package entities

import (
	"gorm.io/gorm"
)

type Midtrans struct {
	gorm.Model
	Transactions []Transactions `gorm:"ForeignKey:MidtransID"`
}
