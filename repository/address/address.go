package address

import (
	"app_airbnb/entities"
	"errors"

	"gorm.io/gorm"
)

type AddressRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *AddressRepository {
	return &AddressRepository{
		database: db,
	}
}

func (tr *AddressRepository) Get() ([]entities.Address, error) {
	arrAddress := []entities.Address{}

	if err := tr.database.Find(&arrAddress).Error; err != nil {
		return nil, err
	}

	return arrAddress, nil
}

func (tr *AddressRepository) GetById(addressId int) (entities.Address, error) {
	arrAddress := entities.Address{}

	if err := tr.database.Preload("Address").Find(&arrAddress, addressId).Error; err != nil {
		return arrAddress, err
	}

	return arrAddress, nil
}

func (tr *AddressRepository) Insert(user_id int, t entities.Address) (entities.Address, error) {
	t.User_ID = int(user_id)

	if err := tr.database.Create(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}

func (tr *AddressRepository) Update(addressId int, user_id int, newAddress entities.Address) (entities.Address, error) {

	res := tr.database.Model(entities.Address{Model: gorm.Model{ID: uint(addressId)}, User_ID: int(user_id)}).Updates(newAddress)

	if res.RowsAffected == 0 {
		return entities.Address{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	address := newAddress

	return address, nil

}

func (tr *AddressRepository) Delete(addressId int) error {

	var address entities.Address

	if err := tr.database.First(&address, addressId).Error; err != nil {
		return err
	}
	tr.database.Delete(&address, addressId)
	return nil

}
