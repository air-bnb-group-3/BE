package entities

import (
	"gorm.io/gorm"
)

type Rooms struct {
	gorm.Model
	CategoryID  uint
	UserID      uint
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Description string   `gorm:"type:text" json:"description"`
	Price       int      `gorm:"type:int(11)" json:"price"`
	Images      []Images `gorm:"ForeignKey:RoomsID"`
	TotalPerson int      `gorm:"type:int(11)" json:"total_person"`
	TotalRooms  int      `gorm:"type:int(11)" json:"total_rooms"`
	SizeBed     string   `gorm:"type:varchar(100)" json:"size_bed"`
	DateStock   string   `gorm:"type:date" json:"date_stock"`
	Booking     []Booking `gorm:"ForeignKey:RoomsID"`
}
