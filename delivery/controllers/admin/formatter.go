package admin

import "app_airbnb/entities"

type AdminCreateRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type AdminCreateResponseFormat struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToAdminCreateResponseFormat(AdminResponse entities.User) AdminCreateResponseFormat {
	return AdminCreateResponseFormat{
		Name:  AdminResponse.Name,
		Email: AdminResponse.Email,
	}
}

type AdminGetResponseFormat struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToAdminGetResponseFormat(AdminResponses []entities.User) []AdminGetResponseFormat {
	AdminGetResponses := make([]AdminGetResponseFormat, len(AdminResponses))
	for i := 0; i < len(AdminResponses); i++ {
		AdminGetResponses[i].Name = AdminResponses[i].Name
		AdminGetResponses[i].Email = AdminResponses[i].Email
	}
	return AdminGetResponses
}

type ResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
}
