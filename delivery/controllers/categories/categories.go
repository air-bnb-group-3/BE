package categories

import (
	"app_airbnb/delivery/controllers/common"
	"app_airbnb/delivery/middlewares"
	"app_airbnb/entities"
	"app_airbnb/repository/categories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoriesController struct {
	repo categories.Categories
}

func New(repository categories.Categories) *CategoriesController {
	return &CategoriesController{
		repo: repository,
	}
}

func (ctrl *CategoriesController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractRoles(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.BadRequest(http.StatusUnauthorized, "access denied", nil))
		}

		newCategory := CategoryCreateRequestFormat{}

		if err := c.Bind(&newCategory); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		NewCategory := entities.Categories{
			City: newCategory.City,
		}

		res, err := ctrl.repo.Insert(NewCategory)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create Category", ToCategoryCreateResponseFormat(res)))
	}
}

func (ctrl *CategoriesController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ctrl.repo.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Category", ToCategoryGetResponseFormat(res)))
	}
}

func (ctrl *CategoriesController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		categoryId, _ := strconv.Atoi(c.Param("categoryid"))

		res, err := ctrl.repo.GetById(uint(categoryId))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Category By ID", ToCategoryByIdGetResponseFormat(res)))
	}
}

func (ctrl *CategoriesController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractRoles(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.BadRequest(http.StatusUnauthorized, "access denied", nil))
		}
		categoryId, _ := strconv.Atoi(c.Param("categoryid"))

		var UpdateCategory = UpdateCategoryRequestFormat{}

		if err := c.Bind(&UpdateCategory); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ctrl.repo.Update(UpdateCategory.ToUpdateCategoryRequestFormat(uint(categoryId)))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "success update category", ToUpdateCategoryResponseFormat(res)))
	}
}

func (ctrl *CategoriesController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		categoryId, _ := strconv.Atoi(c.Param("categoryid"))
		err := ctrl.repo.Delete(uint(categoryId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "success delete category", err))
	}
}


