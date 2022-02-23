package images

import "app_airbnb/entities"

type Images interface {
	Insert(newImage entities.Images) (entities.Images, error)
	GetAll(roomId uint) ([]entities.Images, error)
	GetById(imgId, roomId uint) (entities.Images, error)
	Update(imgId, userId uint, newImage entities.Images) (entities.Images, error)
	Delete(imgId, roomId, userId uint) error
}
