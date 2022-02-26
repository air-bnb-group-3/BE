package images

import "app_airbnb/entities"

type ImageCreateRequestFormat struct {
	RoomsID uint   `json:"rooms_id" form:"rooms_id"`
	Image   string `json:"image" form:"image"`
}

type ImageCreateResponseFormat struct {
	RoomsID uint   `json:"rooms_id"`
	Image   string `json:"image"`
}

func ToImageCreateResponseFormat(ImageResponse entities.Images) ImageCreateResponseFormat {
	return ImageCreateResponseFormat{
		RoomsID: ImageResponse.RoomsID,
		Image:   ImageResponse.Image,
	}
}

type ImageGetResponseFormat struct {
	RoomsID uint   `json:"rooms_id"`
	Image   string `json:"image"`
}

func ToImageGetResponseFormat(ImageResponses []entities.Images) []ImageGetResponseFormat {
	ImageGetResponses := make([]ImageGetResponseFormat, len(ImageResponses))
	for i := 0; i < len(ImageResponses); i++ {
		ImageGetResponses[i].RoomsID = ImageResponses[i].RoomsID
		ImageGetResponses[i].Image = ImageResponses[i].Image
	}
	return ImageGetResponses
}

type ImageGetByIdResponseFormat struct {
	RoomsID uint   `json:"rooms_id"`
	Image   string `json:"image"`
}

func ToImageGetByIdResponseFormat(ImageResponse entities.Images) ImageGetByIdResponseFormat {
	return ImageGetByIdResponseFormat{
		RoomsID: ImageResponse.RoomsID,
		Image:   ImageResponse.Image,
	}
}

type UpdateImageRequestFormat struct {
	Image string `json:"image"`
}

func (Update UpdateImageRequestFormat) ToUpdateImageRequestFormat() entities.Images {
	return entities.Images{
		Image: Update.Image,
	}
}

type UpdateImageResponseFormat struct {
	RoomsID uint   `json:"rooms_id"`
	Image   string `json:"image"`
}

func ToUpdateImageResponseFormat(ImageResponse entities.Images) UpdateImageResponseFormat {
	return UpdateImageResponseFormat{
		RoomsID: ImageResponse.RoomsID,
		Image:   ImageResponse.Image,
	}
}

type CreateImage struct {
	RoomsID int `json:"rooms_id" form:"rooms_id"`
	Url     string
}

type UploadImage struct {
	Url string `json:"image"`
}

type ImageRequest struct {
	Array []UploadImage `json:"array"`
}
