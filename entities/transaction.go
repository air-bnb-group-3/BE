package entities

import "gorm.io/gorm"

type Transactions struct {
	gorm.Model
	Booking          []Booking `gorm:"ForeignKey:TransactionsID"`
	MidtransID       uint
	PaymentMethodsID uint
	Phone            string `gorm:"type:varchar(13); not null" json:"phone"`
	Status           string `gorm:"type:enum('Payed','Canceled')" json:"status"`
}
