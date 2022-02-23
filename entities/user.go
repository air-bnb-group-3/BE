package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(100)" json:"name"`
	Email    string    `gorm:"unique" json:"email"`
	Password string    `json:"-"`
	Roles    bool      `gorm:"type:bool" json:"roles"`
	Rooms    []Rooms   `gorm:"ForeignKey:UserID"`
	Booking  []Booking `gorm:"ForeignKey:UserID"`
}
