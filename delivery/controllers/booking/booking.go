package booking

import (
	"app_airbnb/delivery/controllers/common"
	"app_airbnb/delivery/middlewares"
	"app_airbnb/entities"
	"app_airbnb/repository/booking"
	_mt "app_airbnb/utils/midtrans"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/datatypes"
)

type BookingController struct {
	repo     booking.Booking
	midtrans coreapi.Client
}

func New(repository booking.Booking, midtrans coreapi.Client) *BookingController {
	return &BookingController{
		repo:     repository,
		midtrans: midtrans,
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

		// days := int(Check_out.Sub(Check_in) / 24)

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

func (bc *BookingController) CreatePayment() echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		booking_id, _ := strconv.Atoi(c.Param("booking_id"))
		var payment_method PaymentTypeRequest
		// user := middlewares.ExtractTokenId(c)

		c.Bind(&payment_method)

		if err := v.Struct(payment_method); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(nil, "There is some problem from input", nil))
		}

		var req *coreapi.ChargeReq

		res_booking, err := bc.repo.GetByID(uint(booking_id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "Your booking is not found", nil))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "Your booking is not found", nil))
		}
		switch payment_method.Payment_method {
		case "klikbca":
			req = &coreapi.ChargeReq{
				PaymentType: coreapi.PaymentTypeBCAKlikpay,
				BCAKlikPay: &coreapi.BCAKlikPayDetails{
					Desc: "PAYMENT PROCESS WITH BCAKLIKPAY",
				},
				TransactionDetails: midtrans.TransactionDetails{
					OrderID:  strconv.Itoa(int(res_booking.ID)), /*id booking*/
					GrossAmt: 40000,                             /*		GrossAmt = price * QTY		*/
				},
				Items: &[]midtrans.ItemDetails{
					{Name: strconv.Itoa(int(res_booking.UserID)), Price: 20000, Qty: 2},
					/*
						Price : Price Rooms yang dipilih
						QTY : selesih hari pemesanan (days := int(Check_out.Sub(Check_in) / 24))
					*/
				},
			}
		}
		apiRes, err := _mt.CreateTransaction(bc.midtrans, req)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "Failed to create payment", nil))

		}

		var responseData PaymentResponse

		responseData.OrderID = apiRes.OrderID
		responseData.GrossAmount = apiRes.GrossAmount
		responseData.PaymentType = apiRes.PaymentType
		responseData.Url = apiRes.Actions[1].URL

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success create payment booking", responseData))
	}
}

func (bc *BookingController) CallBack() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request RequestCallBackMidtrans
		order_id, _ := strconv.Atoi(request.Order_id)

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "Failed to create payment", nil))
		}

		res, err := bc.repo.GetByMidtransID(order_id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "internal server eror for get booking by id "+err.Error(), nil))
		}

		switch request.Transaction_status {
		case "settlement":
			bc.repo.Update(res.UserID, uint(order_id), entities.Booking{Status: "Paid"})
		case "failure":
			bc.repo.Update(res.UserID, uint(order_id), entities.Booking{Status: "waiting"})
		case "cancel":
			bc.repo.Update(res.UserID, uint(order_id), entities.Booking{Status: "waiting"})
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success create payment booking", request))

	}
}
