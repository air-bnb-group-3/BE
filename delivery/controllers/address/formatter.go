package address

import (
	"app_airbnb/entities"
)

// =================== Create Address =======================
type InsertAddressRequestFormat struct {
	Street      string `json:"street" form:"street"`
	City        string `json:"city" form:"city"`
	Region      string `json:"region" form:"region"`
	Postal_code string `json:"postal_code" form:"postal_code"`
}

type AddressResponseFormat struct {
	Street      string `json:"street"`
	City        string `json:"city"`
	Region      string `json:"region"`
	Postal_code string `json:"postal_code"`
}

func ToAddressResponseFormat(AddressResponse entities.Address) AddressResponseFormat {
	return AddressResponseFormat{
		Street:   AddressResponse.Street,
		City:     AddressResponse.City,
		Region: AddressResponse.Region,
		Postal_code: AddressResponse.Postal_code,
	}
}

// =================== Update Address =======================
type UpdateAddressRequestFormat struct {
	Street      string `json:"street" form:"street"`
	City        string `json:"city" form:"city"`
	Region      string `json:"region" form:"region"`
	Postal_code string `json:"postal_code" form:"postal_code"`
}

func (UARF UpdateAddressRequestFormat) ToUpdateAddressRequestFormat() entities.Address {
	return entities.Address{
		Street:   UARF.Street,
		City:     UARF.City,
		Region: UARF.Region,
		Postal_code: UARF.Postal_code,
	}
}

type UpdateAddressResponseFormat struct {
	Street      string `json:"street"`
	City        string `json:"city"`
	Region      string `json:"region"`
	Postal_code string `json:"postal_code"`
}

func ToUpdateAddressResponseFormat(AddressResponse entities.Address) UpdateAddressResponseFormat {
	return UpdateAddressResponseFormat{
		Street:   AddressResponse.Street,
		City:     AddressResponse.City,
		Region: AddressResponse.Region,
		Postal_code: AddressResponse.Postal_code,
	}
}