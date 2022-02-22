package admin

import (
	"app_airbnb/delivery/middlewares"
	"app_airbnb/entities"

	"gorm.io/gorm"
)

type AdminRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *AdminRepository {
	return &AdminRepository{
		database: db,
	}
}

func (ur *AdminRepository) Register(newAdmin entities.User) (entities.User, error) {

	newAdmin.Password, _ = middlewares.HashPassword(newAdmin.Password)

	if err := ur.database.Create(&newAdmin).Error; err != nil {
		return newAdmin, err
	}

	return newAdmin, nil
}

func (ur *AdminRepository) GetAll() ([]entities.User, error) {
	arrUser := []entities.User{}

	if err := ur.database.Model(&arrUser).Find(&arrUser).Error; err != nil {
		return nil, err
	}

	return arrUser, nil
}