package admin

import (
	"app_airbnb/delivery/controllers/common"
	"app_airbnb/delivery/middlewares"
	// "app_airbnb/delivery/middlewares"
	"app_airbnb/entities"
	// "app_airbnb/repository/user"
	"app_airbnb/repository/admin"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	repo admin.Admin
}

func New(repository admin.Admin) *AdminController {
	return &AdminController{
		repo: repository,
	}
}

func (ac *AdminController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		newAdmin := AdminCreateRequestFormat{}

		if err := c.Bind(&newAdmin); err != nil || newAdmin.Email == "" || newAdmin.Password == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		NewAdmin := entities.User{
			Name:     newAdmin.Name,
			Email:    newAdmin.Email,
			Password: newAdmin.Password,
			Roles:  true,
		}

		res, err := ac.repo.Register(NewAdmin)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create Admin", ToAdminCreateResponseFormat(res)))
	}
}

func (ac *AdminController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractRoles(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.BadRequest(http.StatusUnauthorized, "invalid input", nil))
		}

		res, err := ac.repo.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All User", ToAdminGetResponseFormat(res)))
	}
}

// func (ac *UserController) GetAll() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// email := middlewares.ExtractTokenAdmin(c)[0]
// 		// password := middlewares.ExtractTokenAdmin(c)[1]

// 		if email != "admin@admin.com" && password != "admin" {
// 			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "invalid input", nil))
// 		}

// 		res, err := ac.repo.GetAll()

// 		if err != nil || email != "admin@admin.com" && password != "admin" {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
// 		}

// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All User", res))
// 	}
// }