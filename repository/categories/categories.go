package categories

import (
	"app_airbnb/entities"
	"errors"

	"gorm.io/gorm"
)

type CategoriesRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CategoriesRepository {
	return &CategoriesRepository{db: db}
}

// ======================== Insert Category =================================
func (repo *CategoriesRepository) Insert(newCategory entities.Categories) (entities.Categories, error) {
	if err := repo.db.Create(&newCategory).Error; err != nil {
		return newCategory, errors.New("gagal memasukkan data category")
	}
	return newCategory, nil
}

// ======================== Get All Category =================================
func (repo *CategoriesRepository) GetAll() ([]entities.Categories, error) {
	categories := []entities.Categories{}
	repo.db.Preload("Rooms").Find(&categories)
	if len(categories) < 1 {
		return nil, errors.New("belum ada category yang terdaftar")
	}
	return categories, nil
}

// ======================== Get Category By ID ==================================
func (repo *CategoriesRepository) GetById(categoryId uint) (entities.Categories, error) {
	category := entities.Categories{}
	if err := repo.db.Preload("Rooms").Preload("").Where("id = ?", categoryId).First(&category).Error; err != nil {
		return category, errors.New("category yang dipilih belum tersedia")
	}
	return category, nil
}

// ======================== Update Category =================================
func (repo *CategoriesRepository) Update(categoryID uint, categoriesUpdate entities.Categories) (entities.Categories, error) {

	res := repo.db.Model(&entities.Categories{Model: gorm.Model{ID: categoryID}}).Updates(categoriesUpdate)
	if res.RowsAffected == 0 {
		return categoriesUpdate, errors.New("tidak ada pemutakhiran pada data category")
	}
	repo.db.First(&categoriesUpdate)
	return categoriesUpdate, nil
}

// ======================== Delete Category =================================
func (repo *CategoriesRepository) Delete(categoryId uint) error {
	categories := entities.Categories{}
	res := repo.db.Delete(&categories, categoryId)
	if res.RowsAffected == 0 {
		return errors.New("tidak ada category yang dihapus")
	}
	return nil
}

// ============================================================================
