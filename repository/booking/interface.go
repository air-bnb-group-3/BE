package booking

import (
	"app_airbnb/entities"

	"gorm.io/datatypes"
)

type Booking interface {
	Create(newBooking entities.Booking) (entities.Booking, error)
	GetByUserID(userId uint) ([]entities.Booking, error)
	Update(bookingId, userId uint, bookingUpdate entities.Booking) (entities.Booking, error)
	GetByID(Id uint) (BookingGetByIdResp, error)
	GetByMidtransID(bookingId int) (entities.Booking, error)
	Delete(user_id int, booking_id int) error
}

type BookingGetByIdResp struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	RoomsID     uint           `json:"rooms_id"`
	Description string         `json:"description"`
	CheckIn     datatypes.Date `json:"check_in"`
	CheckOut    datatypes.Date `json:"check_out"`
	Price       int            `json:"price"`
	UserID      uint           `json:"user_id"`
	Days        int            `json:"days"`
	PriceTotal  int            `json:"price_total"`
	Status      string         `json:"status"`
}
