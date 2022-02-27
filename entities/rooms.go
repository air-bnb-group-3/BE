package entities

import (
	"gorm.io/gorm"
)

type Rooms struct {
	gorm.Model
	CategoryID  uint
	UserID      uint
	Name        string    `gorm:"type:varchar(100)" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Address     string    `gorm:"type:text" json:"address"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Price       int       `gorm:"type:int(11)" json:"price"`
	TotalPerson int       `gorm:"type:int(11)" json:"total_person"`
	TotalRooms  int       `gorm:"type:int(11)" json:"total_rooms"`
	SizeBed     string    `gorm:"type:varchar(100)" json:"size_bed"`
	Images      []Images  `gorm:"ForeignKey:RoomsID"`
	Booking     []Booking `gorm:"ForeignKey:RoomsID"`
}
