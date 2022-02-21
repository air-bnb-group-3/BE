package address

import (
	"app_airbnb/entities"
)

//----------------------------------------------------
//REQUEST FORMAT
//----------------------------------------------------
type InsertAddressRequestFormat struct {
	Street      string `json:"street" form:"street"`
	City        string `json:"city" form:"city"`
	Region      string `json:"region" form:"region"`
	Postal_code string `json:"postal_code" form:"postal_code"`
	User_ID     int    `json:"user_id" form:"user_id"`
}
type UpdateAddressRequestFormat struct {
	Street      string `json:"street" form:"street"`
	City        string `json:"city" form:"city"`
	Region      string `json:"region" form:"region"`
	Postal_code string `json:"postal_code" form:"postal_code"`
}

//-----------------------------------------------------
//RESPONSE FORMAT
//-----------------------------------------------------
type InsertAddressResponseFormat struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    entities.Address `json:"data"`
}

type GetAddresssResponseFormat struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    []entities.Address `json:"data"`
}

type GetAllAddressResponseFormat struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    entities.Address `json:"data"`
}

type UpdateAddressResponseFormat struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    entities.Address `json:"data"`
}

type DeleteAddressResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
