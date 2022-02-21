package owner

import (
	"app_airbnb/delivery/controllers/common"
	// "app_airbnb/delivery/middlewares"
	"app_airbnb/entities"
	// "app_airbnb/repository/user"
	"app_airbnb/repository/owner"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OwnerController struct {
	repo owner.Owner
}

func New(repository owner.Owner) *OwnerController {
	return &OwnerController{
		repo: repository,
	}
}

func (ac *OwnerController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		newOwner := OwnerCreateRequestFormat{}

		if err := c.Bind(&newOwner); err != nil || newOwner.Email == "" || newOwner.Password == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		NewOwner := entities.User{
			Name:     newOwner.Name,
			Email:    newOwner.Email,
			Password: newOwner.Password,
			Roles:  true,
		}

		res, err := ac.repo.Register(NewOwner)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create Owner", ToOwnerCreateResponseFormat(res)))
	}
}