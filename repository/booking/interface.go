package booking

import "app_airbnb/entities"

type Booking interface {
	Create(newBooking entities.Booking) (entities.Booking, error)
	GetByUserID(userId uint) ([]entities.Booking, error)
	Update(bookingId, userId uint, bookingUpdate entities.Booking) (entities.Booking, error)
	GetByID(Id uint) (entities.Booking, error)
	Delete(user_id int, booking_id int) error
}
