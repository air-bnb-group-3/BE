package categories

import (
	"app_airbnb/entities"
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
	ID    int    `json:"id"`
	City  string `json:"city"`
	Rooms []entities.Rooms
}

func ToCategoryGetResponseFormat(CategoryResponses []entities.Categories) []CategoryGetResponseFormat {
	CategoryGetResponses := make([]CategoryGetResponseFormat, len(CategoryResponses))
	for i := 0; i < len(CategoryResponses); i++ {
		CategoryGetResponses[i].ID = int(CategoryResponses[i].ID)
		CategoryGetResponses[i].City = CategoryResponses[i].City
		CategoryGetResponses[i].Rooms = CategoryResponses[i].Rooms

	}
	return CategoryGetResponses
}

type CategoryGetByIdResponseFormat struct {
	ID    int    `json:"id"`
	City  string `json:"city"`
	Rooms []entities.Rooms
}

func ToCategoryByIdGetResponseFormat(CategoryRespon entities.Categories) CategoryGetByIdResponseFormat {
	return CategoryGetByIdResponseFormat{
		ID:    int(CategoryRespon.ID),
		City:  CategoryRespon.City,
		Rooms: CategoryRespon.Rooms,
	}
}

type UpdateCategoryRequestFormat struct {
	City string `json:"city" form:"city"`
}

func (UCRF UpdateCategoryRequestFormat) ToUpdateCategoryRequestFormat() entities.Categories {
	return entities.Categories{
		City: UCRF.City,
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
