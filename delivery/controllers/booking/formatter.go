package booking

import "app_airbnb/entities"

type RequestBooking struct {
	RoomsID  uint `json:"rooms_id" form:"rooms_id"`
	DateTime int  `json:"date_time" form:"date_time"`
}

type ResponseBooking struct {
	ID          uint `json:"id"`
	UserID      uint `json:"user_id"`
	RoomsID     uint `json:"rooms_id"`
	DateTime    int  `json:"date_time"`
	Total_price int  `json:"total_price"`
}

type ResponseGetBooking struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []entities.Booking
}
type ResponseGetBookingByIdBooking struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    entities.Booking
}
