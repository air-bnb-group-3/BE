package booking

import (
	"app_airbnb/delivery/controllers/common"
	"app_airbnb/delivery/middlewares"
	"app_airbnb/entities"
	"app_airbnb/repository/booking"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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
		user_id := int(middlewares.ExtractTokenId(c))
		newBooking := RequestBooking{}

		if err := c.Bind(&newBooking); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(
				http.StatusBadRequest,
				"There is some problem from input",
				nil))
		}

		res, err := bc.repo.Create(user_id, entities.Booking{RoomsID: uint(newBooking.RoomsID)})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(
				http.StatusInternalServerError,
				"error in database process",
				nil,
			))
		}

		return c.JSON(http.StatusCreated, common.Success(
			http.StatusCreated,
			"success to create Booking",
			res,
		))
	}
}

func (bc *BookingController) GetByUserId() echo.HandlerFunc {
	return func(c echo.Context) error {

		user_id := int(middlewares.ExtractTokenId(c))
		res, err := bc.repo.GetByUID(uint(user_id))

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
			res,
		))
	}
}

func (bc *BookingController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		Id, _ := strconv.Atoi(c.Param("booking_id"))

		res, err := bc.repo.GetByID(uint(Id))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Booking By ID", res))
	}
}

func (bc *BookingController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		// id, _ := strconv.Atoi(c.Param("id"))
		user_id := int(middlewares.ExtractTokenId(c))
		upBooking := RequestBooking{}

		if err := c.Bind(&upBooking); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(
				http.StatusBadRequest,
				"error in input cart",
				nil,
			))
		}

		res, err := bc.repo.Update(user_id, entities.Booking{RoomsID: upBooking.RoomsID})

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
			res,
		))
	}
}

func (bc *BookingController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		cart_id, _ := strconv.Atoi(c.Param("id"))

		user_id := int(middlewares.ExtractTokenId(c))

		err := bc.repo.Delete(user_id, cart_id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(
				http.StatusInternalServerError,
				"error in database process",
				nil,
			))
		}

		return c.JSON(http.StatusOK, common.Success(
			http.StatusOK,
			"success to delete booking",
			nil,
		))
	}
}
