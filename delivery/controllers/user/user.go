package user

import (
	"app_airbnb/delivery/controllers/common"
	"app_airbnb/delivery/middlewares"
	"app_airbnb/entities"
	"app_airbnb/repository/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo user.User
}

func New(repository user.User) *UserController {
	return &UserController{
		repo: repository,
	}
}

func (ac *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := InsertUserRequestFormat{}

		if err := c.Bind(&user); err != nil || user.Email == "" || user.Password == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		if user.Email == "admin@admin.com" && user.Password == "admin" {
			user.Status = "admin"
		} else {
			user.Status = "user"
		}

		res, err := ac.repo.Register(entities.User{Name: user.Name, Email: user.Email, Password: user.Password, Status: user.Status})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create User", res))
	}
}

func (ac *UserController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := int(middlewares.ExtractTokenId(c))

		res, err := ac.repo.GetById(userId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "Not Found", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get User", res))
	}
}

func (ac *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := int(middlewares.ExtractTokenId(c))
		var newUser = UpdateUserRequestFormat{}

		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Update(userId, entities.User{Name: newUser.Name, Email: newUser.Email, Status: newUser.Status})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update User", res))
	}
}

func (ac *UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := int(middlewares.ExtractTokenId(c))

		err := ac.repo.Delete(userId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete User", nil))
	}
}

func (ac *UserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := middlewares.ExtractTokenAdmin(c)[0]
		password := middlewares.ExtractTokenAdmin(c)[1]

		if email != "admin@admin.com" && password != "admin" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "invalid input", nil))
		}

		res, err := ac.repo.GetAll()

		if err != nil || email != "admin@admin.com" && password != "admin" {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All User", res))
	}
}
