package rooms

import (
	"app_airbnb/entities"
)

type RoomCreateRequestFormat struct {
	CategoryID  uint    `json:"category_id" form:"category_id"`
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Address     string  `json:"address" form:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Price       int     `json:"price" form:"price"`
	TotalPerson int     `json:"total_person" form:"total_person"`
	TotalRooms  int     `json:"total_rooms" form:"total_rooms"`
	SizeBed     string  `json:"size_bed" form:"size_bed"`
}

func (RCRF RoomCreateRequestFormat) ToRoomEntity(UserID uint) entities.Rooms {
	return entities.Rooms{
		CategoryID:  RCRF.CategoryID,
		Name:        RCRF.Name,
		Description: RCRF.Description,
		Address:     RCRF.Address,
		Latitude:    RCRF.Latitude,
		Longitude:   RCRF.Longitude,
		Price:       RCRF.Price,
		TotalPerson: RCRF.TotalPerson,
		TotalRooms:  RCRF.TotalRooms,
		SizeBed:     RCRF.SizeBed,
		UserID:      UserID,
	}
}

type RoomCreateResponseFormat struct {
	ID          uint    `json:"id"`
	CategoryID  uint    `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Price       int     `json:"price"`
	TotalPerson int     `json:"total_person"`
	TotalRooms  int     `json:"total_rooms"`
	SizeBed     string  `json:"size_bed"`
	Images      []entities.Images
}

func ToRoomCreateResponseFormat(RoomResponse entities.Rooms) RoomCreateResponseFormat {
	return RoomCreateResponseFormat{
		ID:         RoomResponse.ID,
		CategoryID: RoomResponse.CategoryID,
		Name:       RoomResponse.Name,
		// Description: RoomResponse.Description,
		// Address:     RoomResponse.Address,
		// Price:       RoomResponse.Price,
		// TotalPerson: RoomResponse.TotalPerson,
		// TotalRooms:  RoomResponse.TotalRooms,
		// SizeBed:     RoomResponse.SizeBed,
		// Images:      RoomResponse.Images,
	}
}

type RoomGetResponseFormat struct {
	ID          uint    `json:"id"`
	CategoryID  uint    `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Price       int     `json:"price"`
	TotalPerson int     `json:"total_person"`
	TotalRooms  int     `json:"total_rooms"`
	SizeBed     string  `json:"size_bed"`
}

func ToRoomGetResponseFormat(RoomResponses []entities.Rooms) []RoomGetResponseFormat {
	RoomGetResponses := make([]RoomGetResponseFormat, len(RoomResponses))
	for i := 0; i < len(RoomResponses); i++ {
		RoomGetResponses[i].ID = RoomResponses[i].ID
		RoomGetResponses[i].CategoryID = RoomResponses[i].CategoryID
		RoomGetResponses[i].Name = RoomResponses[i].Name
		RoomGetResponses[i].Description = RoomResponses[i].Description
		RoomGetResponses[i].Address = RoomResponses[i].Address
		RoomGetResponses[i].Latitude = RoomResponses[i].Latitude
		RoomGetResponses[i].Longitude = RoomResponses[i].Longitude
		RoomGetResponses[i].Price = RoomResponses[i].Price
		RoomGetResponses[i].TotalPerson = RoomResponses[i].TotalPerson
		RoomGetResponses[i].TotalRooms = RoomResponses[i].TotalRooms
		RoomGetResponses[i].SizeBed = RoomResponses[i].SizeBed
	}
	return RoomGetResponses
}

type RoomGetByIdResponseFormat struct {
	ID          uint    `json:"id"`
	CategoryID  uint    `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Address     string  `json:"address"`
	Price       int     `json:"price"`
	TotalPerson int     `json:"total_person"`
	TotalRooms  int     `json:"total_rooms"`
	SizeBed     string  `json:"size_bed"`
	Images      []entities.Images
}

func ToRoomByIdGetResponseFormat(RoomResponse entities.Rooms) RoomGetByIdResponseFormat {
	return RoomGetByIdResponseFormat{
		ID:          RoomResponse.ID,
		CategoryID:  RoomResponse.CategoryID,
		Name:        RoomResponse.Name,
		Description: RoomResponse.Description,
		Address:     RoomResponse.Address,
		Latitude:    RoomResponse.Latitude,
		Longitude:   RoomResponse.Longitude,
		Price:       RoomResponse.Price,
		TotalPerson: RoomResponse.TotalPerson,
		TotalRooms:  RoomResponse.TotalRooms,
		SizeBed:     RoomResponse.SizeBed,
		Images:      RoomResponse.Images,
	}
}

type UpdateRoomRequestFormat struct {
	CategoryID  uint    `json:"category_id" form:"category_id"`
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Address     string  `json:"address" form:"address"`
	Latitude    float64 `json:"latitude" form:"latitutde"`
	Longitude   float64 `json:"longitude" form:"longitude"`
	Price       int     `json:"price" form:"price"`
	TotalPerson int     `json:"total_person" form:"total_person"`
	TotalRooms  int     `json:"total_rooms" form:"total_rooms"`
	SizeBed     string  `json:"size_bed" form:"size_bed"`
}

func (URRF UpdateRoomRequestFormat) ToUpdateRoomRequestFormat() entities.Rooms {
	return entities.Rooms{
		CategoryID:  URRF.CategoryID,
		Name:        URRF.Name,
		Description: URRF.Description,
		Address:     URRF.Address,
		Latitude:    URRF.Latitude,
		Longitude:   URRF.Longitude,
		Price:       URRF.Price,
		TotalPerson: URRF.TotalPerson,
		TotalRooms:  URRF.TotalRooms,
		SizeBed:     URRF.SizeBed,
	}
}

type UpdateRoomResponseFormat struct {
	ID          uint    `json:"id"`
	CategoryID  uint    `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Price       int     `json:"price"`
	TotalPerson int     `json:"total_person"`
	TotalRooms  int     `json:"total_rooms"`
	SizeBed     string  `json:"size_bed"`
}

func ToUpdateRoomResponseFormat(RoomResponse entities.Rooms) UpdateRoomResponseFormat {
	return UpdateRoomResponseFormat{
		ID:          RoomResponse.ID,
		CategoryID:  RoomResponse.CategoryID,
		Name:        RoomResponse.Name,
		Description: RoomResponse.Description,
		Address:     RoomResponse.Address,
		Latitude:    RoomResponse.Latitude,
		Longitude:   RoomResponse.Longitude,
		Price:       RoomResponse.Price,
		TotalPerson: RoomResponse.TotalPerson,
		TotalRooms:  RoomResponse.TotalRooms,
		SizeBed:     RoomResponse.SizeBed,
	}
}
