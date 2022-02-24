package images

import "app_airbnb/entities"

type Images interface {
	Insert(newImage entities.Images) (entities.Images, error)
	GetAll() ([]entities.Images, error)
	GetById(imgId uint) (entities.Images, error)
	Update(imgId, userId uint, newImage entities.Images) (entities.Images, error)
	Delete(imgId, userId uint) error
}
