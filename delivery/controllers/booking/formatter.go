package booking

import (
	"app_airbnb/entities"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type BookingCreateRequestFormat struct {
	RoomsID        uint   `json:"rooms_id" form:"rooms_id"`
	CheckIn        string `json:"check_in" form:"check_in"`
	CheckOut       string `json:"check_out" form:"check_out"`
	PaymentMethods string `json:"payment_methods" form:"payment_methods"`
	Status         string `json:"status" form:"status"`
}

func (Booking BookingCreateRequestFormat) ToBookingEntity(CheckIn, CheckOut datatypes.Date, UserID uint) entities.Booking {
	return entities.Booking{
		RoomsID:        Booking.RoomsID,
		CheckIn:        CheckIn,
		CheckOut:       CheckOut,
		PaymentMethods: Booking.PaymentMethods,
		Status:         Booking.Status,
		UserID:         UserID,
	}
}

type BookingCreateResponseFormat struct {
	ID             uint
	RoomsID        uint           `json:"rooms_id"`
	CheckIn        datatypes.Date `json:"check_in"`
	CheckOut       datatypes.Date `json:"check_out"`
	PaymentMethods string         `json:"payment_methods"`
	Status         string         `json:"status"`
}

func ToBookingCreateResponseFormat(BookingResponse entities.Booking) BookingCreateResponseFormat {
	return BookingCreateResponseFormat{
		ID:             BookingResponse.ID,
		RoomsID:        BookingResponse.RoomsID,
		CheckIn:        BookingResponse.CheckIn,
		CheckOut:       BookingResponse.CheckOut,
		PaymentMethods: BookingResponse.PaymentMethods,
		Status:         BookingResponse.Status,
	}
}

type BookingGetResponseFormat struct {
	ID             uint
	RoomsID        uint           `json:"rooms_id"`
	CheckIn        datatypes.Date `json:"check_in"`
	CheckOut       datatypes.Date `json:"check_out"`
	PaymentMethods string         `json:"payment_methods"`
	Status         string         `json:"status"`
	// Days           int            `json:"days"`
	// Price          int            `json:"price"`
}

func ToBookingGetResponseFormat(BookingResponses []entities.Booking) []BookingGetResponseFormat {
	BookingGetResponses := make([]BookingGetResponseFormat, len(BookingResponses))
	for i := 0; i < len(BookingResponses); i++ {
		BookingGetResponses[i].ID = BookingResponses[i].ID
		BookingGetResponses[i].RoomsID = BookingResponses[i].RoomsID
		BookingGetResponses[i].CheckIn = BookingResponses[i].CheckIn
		BookingGetResponses[i].CheckOut = BookingResponses[i].CheckOut
		BookingGetResponses[i].PaymentMethods = BookingResponses[i].PaymentMethods
		BookingGetResponses[i].Status = BookingResponses[i].Status
		// BookingGetResponses[i].Days = BookingGetResponses[i].Days
		// BookingGetResponses[i].Price = BookingGetResponses[i].Price
	}
	return BookingGetResponses
}

type UpdateBookingRequestFormat struct {
	RoomsID  uint   `json:"rooms_id" form:"rooms_id"`
	CheckIn  string `json:"check_in" form:"check_in"`
	CheckOut string `json:"check_out" form:"check_out"`
}

func (Booking UpdateBookingRequestFormat) ToUpdateBookingRequestFormat(bookingId uint, CheckIn, CheckOut datatypes.Date) entities.Booking {
	return entities.Booking{
		Model:    gorm.Model{ID: bookingId},
		RoomsID:  Booking.RoomsID,
		CheckIn:  CheckIn,
		CheckOut: CheckOut,
	}
}

type UpdateBookingResponseFormat struct {
	ID             uint
	RoomsID        uint           `json:"rooms_id"`
	CheckIn        datatypes.Date `json:"check_in"`
	CheckOut       datatypes.Date `json:"check_out"`
	PaymentMethods string         `json:"payment_methods"`
	Status         string         `json:"status"`
}

func ToUpdateBookingResponseFormat(BookingResponse entities.Booking) UpdateBookingResponseFormat {
	return UpdateBookingResponseFormat{
		ID:             BookingResponse.ID,
		RoomsID:        BookingResponse.RoomsID,
		CheckIn:        BookingResponse.CheckIn,
		CheckOut:       BookingResponse.CheckOut,
		PaymentMethods: BookingResponse.PaymentMethods,
		Status:         BookingResponse.Status,
	}
}

type PaymentTypeRequest struct {
	Payment_method string `json:"payment_method" validate:"required"`
}

type PaymentResponse struct {
	OrderID     string `json:"order_id"`
	GrossAmount string `json:"gross_amount"`
	PaymentType string `json:"payment_type"`
	Url         string `json:"url"`
}

type RequestCallBackMidtrans struct {
	Transaction_time   string `json:"transaction_time"`
	Transaction_status string `json:"transaction_status"`
	Transaction_id     string `json:"transaction_id"`
	Status_message     string `json:"status_message"`
	Status_code        string `json:"status_code"`
	Signature_key      string `json:"signature_key"`
	Settlement_time    string `json:"settlement_time"`
	Payment_type       string `json:"payment_type"`
	OrderID            string `json:"order_id"`
	Merchant_id        string `json:"merchant_id"`
	Gross_amount       string `json:"gross_amount"`
	Fraud_status       string `json:"fraud_status"`
	Currency           string `json:"currency"`
}
