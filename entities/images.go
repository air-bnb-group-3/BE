package entities

import (
	"gorm.io/gorm"
)

type Images struct {
	gorm.Model
	RoomsID uint
	Image   string `gorm:"type: text" json:"image"`
}
