package rooms

import "app_airbnb/entities"

type Rooms interface {
	Insert(newRooms entities.Rooms) (entities.Rooms, error)
	GetAll() ([]entities.Rooms, error)
	GetById(roomId uint) (entities.Rooms, error)
	Update(roomId , userId uint, newRooms entities.Rooms) (entities.Rooms, error)
	Delete(roomId, userId uint) error
}