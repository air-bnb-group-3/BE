package admin

import "app_airbnb/entities"

type Admin interface {
	Register(newAdmin entities.User) (entities.User, error)
	GetAll() ([]entities.User, error)
	// GetById(userId int) (entities.User, error)
	// Update(userId int, newUser entities.User) (entities.User, error)
	// Delete(userId int) error
}
