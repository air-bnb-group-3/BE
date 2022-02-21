package user

import (
	"app_airbnb/entities"
	// "gorm.io/gorm"
)

// =================== Create User =======================
type CreateUserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type CreateUserResponseFormat struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func ToCreateUserResponseFormat(UserResponse entities.User) CreateUserResponseFormat {
	return CreateUserResponseFormat{
		Name:     UserResponse.Name,
		Email:    UserResponse.Email,
	}
}

// =================== Update User =======================
type UpdateUserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// func (UURF UpdateUserRequestFormat) ToUpdateUserRequestFormat(ID uint) entities.User {
// 	return entities.User{
// 		Model:    gorm.Model{ID: ID},
// 		Name:     UURF.Name,
// 		Email:    UURF.Email,
// 		Password: UURF.Password,
// 	}
// }

type UpdateUserResponseFormat struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func ToUpdateUserResponseFormat(UserResponse entities.User) UpdateUserResponseFormat {
	return UpdateUserResponseFormat{
		Name:     UserResponse.Name,
		Email:    UserResponse.Email,
	}
}

type GetUserByIdResponseFormat struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func ToGetUserByIdResponseFormat(UserResponse entities.User) GetUserByIdResponseFormat {
	return GetUserByIdResponseFormat{
		Name:     UserResponse.Name,
		Email:    UserResponse.Email,
	}
}