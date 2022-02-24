package images

import (
	"app_airbnb/delivery/controllers/common"
	s3 "app_airbnb/utils/aws_S3"
	"strconv"

	// "app_airbnb/delivery/middlewares"

	"app_airbnb/repository/images"
	"net/http"

	// "strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ImagesController struct {
	repo images.Images
	conn *session.Session
}

func New(repository images.Images, aws *session.Session) *ImagesController {
	return &ImagesController{
		repo: repository,
		conn: aws,
	}
}

func (img *ImagesController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		newImage := CreateImage{}
		newImage.RoomsID, _ = strconv.Atoi(c.FormValue("rooms_id"))

		file, err := c.FormFile("file")
		if err != nil {
			log.Info(err)
		}

		fileName := s3.DoUpload(img.conn, *file)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Insert Image", map[string]interface{}{
			"link": fileName,
		}))
	}
}

// func (img *ImagesController) GetAll() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		res, err := img.repo.GetAll()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
// 		}
// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Image", ToImageGetResponseFormat(res)))
// 	}
// }

// func (img *ImagesController) GetById() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		imgId, _ := strconv.Atoi(c.Param("imageid"))

// 		res, err := img.repo.GetById(uint(imgId))
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
// 		}
// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Image By ID", ToImageGetByIdResponseFormat(res)))
// 	}
// }

// func (img *ImagesController) Update() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		UserID := middlewares.ExtractTokenId(c)
// 		imgId, _ := strconv.Atoi(c.Param("imageid"))

// 		var UpdateImage = UpdateImageRequestFormat{}

// 		if err := c.Bind(&UpdateImage); err != nil {
// 			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
// 		}

// 		res, err := img.repo.Update(uint(imgId), uint(UserID), UpdateImage.ToUpdateImageRequestFormat())

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
// 		}
// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "success update image", ToUpdateImageResponseFormat(res)))
// 	}
// }

// func (img *ImagesController) Delete() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		UserID := middlewares.ExtractTokenId(c)
// 		imgId, _ := strconv.Atoi(c.Param("imageid"))
// 		err := img.repo.Delete(uint(imgId), uint(UserID))
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
// 		}
// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "success delete image", err))
// 	}
// }
