package rooms

import (
	"app_airbnb/entities"
)

type RoomCreateRequestFormat struct {
	CategoryID  uint   `json:"category_id" form:"category_id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	TotalPerson int    `json:"total_person" form:"total_person"`
	TotalRooms  int    `json:"total_rooms" form:"total_rooms"`
	SizeBed     string `json:"size_bed" form:"size_bed"`
	DateStock   string `json:"date_stock" form:"date_stock"`
}

func (RCRF RoomCreateRequestFormat) ToRoomEntity(UserID uint) entities.Rooms {
	return entities.Rooms{
		CategoryID:  RCRF.CategoryID,
		Name:        RCRF.Name,
		Description: RCRF.Description,
		Price:       RCRF.Price,
		TotalPerson: RCRF.TotalPerson,
		TotalRooms:  RCRF.TotalRooms,
		SizeBed:     RCRF.SizeBed,
		UserID:      UserID,
	}
}

type RoomCreateResponseFormat struct {
	ID          uint   `json:"id"`
	CategoryID  uint   `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	TotalPerson int    `json:"total_person"`
	TotalRooms  int    `json:"total_rooms"`
	SizeBed     string `json:"size_bed"`
	DateStock   string `json:"date_stock"`
}

func ToRoomCreateResponseFormat(RoomResponse entities.Rooms) RoomCreateResponseFormat {
	return RoomCreateResponseFormat{
		ID:          RoomResponse.ID,
		CategoryID:  RoomResponse.CategoryID,
		Name:        RoomResponse.Name,
		Description: RoomResponse.Description,
		Price:       RoomResponse.Price,
		TotalPerson: RoomResponse.TotalPerson,
		TotalRooms:  RoomResponse.TotalRooms,
		SizeBed:     RoomResponse.SizeBed,
	}
}

type RoomGetResponseFormat struct {
	CategoryID  uint   `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	TotalPerson int    `json:"total_person"`
	TotalRooms  int    `json:"total_rooms"`
	SizeBed     string `json:"size_bed"`
	DateStock   string `json:"date_stock"`
}

func ToRoomGetResponseFormat(RoomResponses []entities.Rooms) []RoomGetResponseFormat {
	RoomGetResponses := make([]RoomGetResponseFormat, len(RoomResponses))
	for i := 0; i < len(RoomResponses); i++ {
		RoomGetResponses[i].CategoryID = RoomResponses[i].CategoryID
		RoomGetResponses[i].Name = RoomResponses[i].Name
		RoomGetResponses[i].Description = RoomResponses[i].Description
		RoomGetResponses[i].Price = RoomResponses[i].Price
		RoomGetResponses[i].TotalPerson = RoomResponses[i].TotalPerson
		RoomGetResponses[i].TotalRooms = RoomResponses[i].TotalRooms
		RoomGetResponses[i].SizeBed = RoomResponses[i].SizeBed
	}
	return RoomGetResponses
}

type RoomGetByIdResponseFormat struct {
	CategoryID  uint   `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	TotalPerson int    `json:"total_person"`
	TotalRooms  int    `json:"total_rooms"`
	SizeBed     string `json:"size_bed"`
	DateStock   string `json:"date_stock"`
}

func ToRoomByIdGetResponseFormat(RoomResponse entities.Rooms) RoomGetByIdResponseFormat {
	return RoomGetByIdResponseFormat{
		CategoryID:  RoomResponse.CategoryID,
		Name:        RoomResponse.Name,
		Description: RoomResponse.Description,
		Price:       RoomResponse.Price,
		TotalPerson: RoomResponse.TotalPerson,
		TotalRooms:  RoomResponse.TotalRooms,
		SizeBed:     RoomResponse.SizeBed,
	}
}

type UpdateRoomRequestFormat struct {
	CategoryID  uint   `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	TotalPerson int    `json:"total_person"`
	TotalRooms  int    `json:"total_rooms"`
	SizeBed     string `json:"size_bed"`
	DateStock   string `json:"date_stock"`
}

func (URRF UpdateRoomRequestFormat) ToUpdateRoomRequestFormat() entities.Rooms {
	return entities.Rooms{
		CategoryID:  URRF.CategoryID,
		Name:        URRF.Name,
		Description: URRF.Description,
		Price:       URRF.Price,
		TotalPerson: URRF.TotalPerson,
		TotalRooms:  URRF.TotalRooms,
		SizeBed:     URRF.SizeBed,
	}
}

type UpdateRoomResponseFormat struct {
	CategoryID  uint   `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	TotalPerson int    `json:"total_person"`
	TotalRooms  int    `json:"total_rooms"`
	SizeBed     string `json:"size_bed"`
	DateStock   string `json:"date_stock"`
}

func ToUpdateRoomResponseFormat(RoomResponse entities.Rooms) UpdateRoomResponseFormat {
	return UpdateRoomResponseFormat{
		CategoryID:  RoomResponse.CategoryID,
		Name:        RoomResponse.Name,
		Description: RoomResponse.Description,
		Price:       RoomResponse.Price,
		TotalPerson: RoomResponse.TotalPerson,
		TotalRooms:  RoomResponse.TotalRooms,
		SizeBed:     RoomResponse.SizeBed,
	}
}
