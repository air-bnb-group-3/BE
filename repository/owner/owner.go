package owner

import (
	"app_airbnb/delivery/middlewares"
	"app_airbnb/entities"

	"gorm.io/gorm"
)

type OwnerRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *OwnerRepository {
	return &OwnerRepository{
		database: db,
	}
}

func (ur *OwnerRepository) Register(newOwner entities.User) (entities.User, error) {

	newOwner.Password, _ = middlewares.HashPassword(newOwner.Password)

	if err := ur.database.Create(&newOwner).Error; err != nil {
		return newOwner, err
	}

	return newOwner, nil
}