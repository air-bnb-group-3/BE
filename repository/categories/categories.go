package categories

import (
	"errors"
	"app_airbnb/entities"

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
	repo.db.Find(&categories)
	if len(categories) < 1 {
		return nil, errors.New("belum ada category yang terdaftar")
	}
	return categories, nil
}

// ======================== Get Category By ID ==================================
func (repo *CategoriesRepository) GetByID(categoryId uint) (entities.Categories, error) {
	category := entities.Categories{}
	if err := repo.db.Model(&category).Where("id = ?", categoryId).First(&category).Error; err != nil {
		return category, errors.New("category yang dipilih belum tersedia")
	}
	return category, nil
}

// ======================== Update Category =================================
func (repo *CategoriesRepository) Update(categoriesUpdate entities.Categories) (entities.Categories, error) {
	categories := entities.Categories{}
	res := repo.db.Model(&categories).Updates(categoriesUpdate)
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
