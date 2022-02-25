package images

import (
	"app_airbnb/entities"
	"errors"

	"gorm.io/gorm"
)

type ImagesRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ImagesRepository {
	return &ImagesRepository{db: db}
}

// ======================== Insert Images =================================
func (repo *ImagesRepository) Insert(newImage entities.Images) (entities.Images, error) {
	if err := repo.db.Create(&newImage).Error; err != nil {
		return newImage, errors.New("gagal memasukkan data image")
	}
	return newImage, nil
}

// // ======================== Get All Images =================================
// func (repo *ImagesRepository) GetAll() ([]entities.Images, error) {
// 	images := []entities.Images{}
// 	repo.db.Model(&images).Find(&images)

// 	// repo.db.Find(&images)
// 	if len(images) < 1 {
// 		return nil, errors.New("belum ada image yang terdaftar")
// 	}
// 	return images, nil
// }

// // ======================== Get Images By ID ==================================
// func (repo *ImagesRepository) GetByID(imageId uint) (entities.Images, error) {
// 	image := entities.Images{}
// 	if err := repo.db.Model(&image).Where("id = ?", imageId).First(&image).Error; err != nil {
// 		return image, errors.New("image yang dipilih belum tersedia")
// 	}
// 	return image, nil
// }

// // ======================== Update Images =================================
// func (repo *ImagesRepository) Update(imageId, userId uint, imagesUpdate entities.Images) (entities.Images, error) {
// 	images := entities.Images{}
// 	res := repo.db.Model(&images).Where("id = ? AND user_id = ?", imageId, userId).Updates(imagesUpdate)

// 	if res.RowsAffected == 0 {
// 		return imagesUpdate, errors.New("tidak ada pemutakhiran pada data image")
// 	}
// 	repo.db.First(&imagesUpdate)
// 	return imagesUpdate, nil
// }

// // ======================== Delete Images =================================
// func (repo *ImagesRepository) Delete(imageId, userId uint) error {
// 	images := entities.Images{}

// 	res := repo.db.Model(&images).Where("id = ? AND user_id = ?", imageId, userId).Delete(&images)
// 	if res.RowsAffected == 0 {
// 		return errors.New("tidak ada image yang dihapus")
// 	}
// 	return nil
// }

// // ============================================================================
