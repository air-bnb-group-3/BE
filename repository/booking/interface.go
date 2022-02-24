package booking

import "app_airbnb/entities"

type Booking interface {
	Create(user_id int, newBook entities.Booking) (entities.Booking, error)
	GetByUID(userId uint) ([]entities.Rooms, error)
	GetByID(Id uint) (entities.Rooms, error)
	Update(userId int, bookingUpdate entities.Booking) (entities.Booking, error)
	Delete(user_id int, booking_id int) error
}
