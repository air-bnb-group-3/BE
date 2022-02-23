package categories

import (
	"app_airbnb/entities"

	"gorm.io/gorm"
)

type CategoryCreateRequestFormat struct {
	City string `json:"city" form:"city"`
}

type CategoryCreateResponseFormat struct {
	City string `json:"city"`
}

func ToCategoryCreateResponseFormat(CategoryResponse entities.Categories) CategoryCreateResponseFormat {
	return CategoryCreateResponseFormat{
		City: CategoryResponse.City,
	}
}

type CategoryGetResponseFormat struct {
	City string `json:"city"`
}

func ToCategoryGetResponseFormat(CategoryResponses []entities.Categories) []CategoryGetResponseFormat {
	CategoryGetResponses := make([]CategoryGetResponseFormat, len(CategoryResponses))
	for i := 0; i < len(CategoryResponses); i++ {
		CategoryGetResponses[i].City = CategoryResponses[i].City
	}
	return CategoryGetResponses
}

type CategoryGetByIdResponseFormat struct {
	City string `json:"city"`
}

func ToCategoryByIdGetResponseFormat(CategoryRespon entities.Categories) CategoryGetByIdResponseFormat {
	return CategoryGetByIdResponseFormat{
		City: CategoryRespon.City,
	}
}

type UpdateCategoryRequestFormat struct {
	City string `json:"city" form:"city"`
}

func (UCRF UpdateCategoryRequestFormat) ToUpdateCategoryRequestFormat(CategoryID uint) entities.Categories {
	return entities.Categories{
		Model: gorm.Model{ID: CategoryID},
		City:  UCRF.City,
	}
}

type UpdateCategoryResponseFormat struct {
	City string `json:"city"`
}

func ToUpdateCategoryResponseFormat(CategoryResponse entities.Categories) UpdateCategoryResponseFormat {
	return UpdateCategoryResponseFormat{
		City: CategoryResponse.City,
	}
}
