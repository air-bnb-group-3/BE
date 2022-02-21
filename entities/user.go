package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `gorm:"type:varchar(100)" json:"name"`
	Email    string  `gorm:"unique" json:"email"`
	Password string  `json:"-"`
	Status   string  `gorm:"type:varchar(50)" json:"status"`
	Address  Address `gorm:"ForeignKey:User_ID"`
}
