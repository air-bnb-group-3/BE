package owner

import "app_airbnb/entities"

type OwnerCreateRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type OwnerCreateResponseFormat struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func ToOwnerCreateResponseFormat(OwnerResponse entities.User) OwnerCreateResponseFormat {
	return OwnerCreateResponseFormat{
		Name:     OwnerResponse.Name,
		Email:    OwnerResponse.Email,
	}
}

// type OwnerGetResponseFormat struct {
// 	Name     string `json:"name"`
// 	Email    string `json:"email"`
// }

// func ToOwnerGetResponseFormat(OwnerResponses []entities.User) []OwnerGetResponseFormat {
// 	OwnerGetResponses := make([]OwnerGetResponseFormat, len(OwnerResponses))
// 	for i := 0; i < len(OwnerResponses); i++ {
// 		OwnerGetResponses[i].Name = OwnerResponses[i].Name
// 		OwnerGetResponses[i].Email = OwnerResponses[i].Email
// 	}
// 	return OwnerGetResponses
// }
