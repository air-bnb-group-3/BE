package address

import "app_airbnb/entities"

type Address interface {
	Get() ([]entities.Address, error)
	GetById(addressId int) (entities.Address, error)
	Insert(user_id int, t entities.Address) (entities.Address, error)
	Update(addressId int, user_id int, newAddress entities.Address) (entities.Address, error)
	Delete(addressId int) error
}
