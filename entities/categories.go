package entities

import (
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	City  string  `gorm:"type:varchar(100)" json:"city"`
	Rooms []Rooms `gorm:"ForeignKey:CategoryID"`
}
