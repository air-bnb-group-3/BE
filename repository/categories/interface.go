package categories

import "app_airbnb/entities"

type Categories interface {
	Insert(newCategory entities.Categories) (entities.Categories, error)
	GetAll() ([]entities.Categories, error)
	GetById(categoryId uint) (entities.Categories, error)
	Update(categoryID uint, updateCategory entities.Categories) (entities.Categories, error)
	Delete(categoryId uint) error
}