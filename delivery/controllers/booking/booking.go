package booking

import (
	"app_airbnb/delivery/controllers/common"
	"app_airbnb/delivery/middlewares"
	"app_airbnb/repository/booking"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

type BookingController struct {
	repo booking.Booking
}

func New(repository booking.Booking) *BookingController {
	return &BookingController{
		repo: repository,
	}
}

func (bc *BookingController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := int(middlewares.ExtractTokenId(c))
		newBooking := BookingCreateRequestFormat{}

		if err := c.Bind(&newBooking); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(
				http.StatusBadRequest,
				"There is some problem from input",
				nil))
		}

		layoutFormat := "2006-01-02 15:04:05"
		Check_in, _ := time.Parse(layoutFormat, newBooking.CheckIn)
		Check_out, _ := time.Parse(layoutFormat, newBooking.CheckOut)

		res, err := bc.repo.Create(newBooking.ToBookingEntity(datatypes.Date(Check_in), datatypes.Date(Check_out), uint(UserID)))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(
				http.StatusInternalServerError,
				"There is some error in server ",
				nil,
			))
		}

		return c.JSON(http.StatusCreated, common.Success(
			http.StatusCreated,
			"success to create Booking",
			ToBookingCreateResponseFormat(res),
		))
	}
}

func (bc *BookingController) GetByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {

		user_id := int(middlewares.ExtractTokenId(c))
		res, err := bc.repo.GetByUserID(uint(user_id))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(
				http.StatusInternalServerError,
				"error in database process",
				nil,
			))
		}

		return c.JSON(http.StatusOK, common.Success(
			http.StatusOK,
			"Success Get Booking by user ID",
			ToBookingGetResponseFormat(res),
		))
	}
}

func (bc *BookingController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		bookingId, _ := strconv.Atoi(c.Param("bookingid"))
		user_id := int(middlewares.ExtractTokenId(c))
		upBooking := UpdateBookingRequestFormat{}

		if err := c.Bind(&upBooking); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(
				http.StatusBadRequest,
				"error in input",
				nil,
			))
		}

		layoutFormat := "2006-01-02 15:04:05"
		Check_in, _ := time.Parse(layoutFormat, upBooking.CheckIn)
		Check_out, _ := time.Parse(layoutFormat, upBooking.CheckOut)
		res, err := bc.repo.Update(uint(bookingId), uint(user_id), upBooking.ToUpdateBookingRequestFormat(uint(bookingId), datatypes.Date(Check_in), datatypes.Date(Check_out)))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(
				http.StatusInternalServerError,
				"error in database process",
				nil,
			))
		}

		return c.JSON(http.StatusOK, common.Success(
			http.StatusOK,
			"success to update booking",
			ToUpdateBookingResponseFormat(res),
		))
	}
}

// func (bc *BookingController) GetById() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		Id, _ := strconv.Atoi(c.Param("booking_id"))

// 		res, err := bc.repo.GetByID(uint(Id))

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
// 		}

// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Booking By ID", res))
// 	}
// }

// func (bc *BookingController) Delete() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		cart_id, _ := strconv.Atoi(c.Param("id"))

// 		user_id := int(middlewares.ExtractTokenId(c))

// 		err := bc.repo.Delete(user_id, cart_id)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(
// 				http.StatusInternalServerError,
// 				"error in database process",
// 				nil,
// 			))
// 		}

// 		return c.JSON(http.StatusOK, common.Success(
// 			http.StatusOK,
// 			"success to delete booking",
// 			nil,
// 		))
// 	}
// }
