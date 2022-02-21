package address

import (
	"app_airbnb/delivery/controllers/common"
	"app_airbnb/delivery/middlewares"
	"app_airbnb/entities"
	"app_airbnb/repository/address"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AddressController struct {
	repo address.Address
}

func New(repository address.Address) *AddressController {
	return &AddressController{
		repo: repository,
	}
}

func (ad *AddressController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ad.repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Address", res))
	}
}

func (ad *AddressController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		addressId, _ := strconv.Atoi(c.Param("id"))

		res, err := ad.repo.GetById(addressId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "Not Found", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Address", res))
	}
}

func (ad *AddressController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := int(middlewares.ExtractTokenId(c))
		address := InsertAddressRequestFormat{}

		if err := c.Bind(&address); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ad.repo.Insert(userId, entities.Address{Street: address.Street, City: address.City, Region: address.Region, Postal_code: address.Postal_code, User_ID: userId})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create Address", res))
	}
}

func (ad *AddressController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		addressId, _ := strconv.Atoi(c.Param("id"))
		user_id := int(middlewares.ExtractTokenId(c))
		newAddress := UpdateAddressRequestFormat{}

		if err := c.Bind(&newAddress); err != nil || newAddress.City == "" && newAddress.Street == "" && newAddress.Region == "" && newAddress.Postal_code == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ad.repo.Update(addressId, user_id, entities.Address{Street: newAddress.Street, City: newAddress.City, Region: newAddress.Region, Postal_code: newAddress.Postal_code})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update Address", res))
	}
}

func (ad *AddressController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		addressId, _ := strconv.Atoi(c.Param("id"))

		err := ad.repo.Delete(addressId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete Address", nil))
	}
}
