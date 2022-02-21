package entities

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street      string `gorm:"type:varchar(150)" json:"street"`
	City        string `gorm:"type:varchar(100)" json:"city"`
	Region      string `gorm:"type:varchar(50)" json:"region"`
	Postal_code string `gorm:"type:varchar(50)" json:"postal_code"`
	User_ID     int    `gorm:"column:user_id" json:"user_id"`
}
