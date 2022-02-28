package rooms

import (
	"app_airbnb/delivery/controllers/common"
	"app_airbnb/delivery/middlewares"
	"app_airbnb/entities"
	"app_airbnb/repository/images"
	"app_airbnb/repository/rooms"
	s3 "app_airbnb/utils/aws_S3"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoomsController struct {
	repo      rooms.Rooms
	repoImage images.Images
}

func New(repository rooms.Rooms, repoImage images.Images) *RoomsController {
	return &RoomsController{
		repo:      repository,
		repoImage: repoImage,
	}
}

func (ctrl *RoomsController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenId(c)
		newRoom := RoomCreateRequestFormat{}

		if err := c.Bind(&newRoom); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ctrl.repo.Insert(newRoom.ToRoomEntity(uint(UserID)))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		form, errM := c.MultipartForm()
		if errM != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "Error Multiform Input", nil))
		}
		files := form.File["files"]

		for _, file := range files {
			newImage := entities.Images{}
			src, _ := file.Open()
			s := s3.InitS3("AKIAVOMUO3KKNSP4RXWR", "o3T3ozzKzrdIfiDTPMVFMgP7NWfpFm75hxtX2Cww", "ap-southeast-1")
			filename := s3.Upload(s, src, file)
			newImage.Image = "https://airbnb-app.s3.ap-southeast-1.amazonaws.com/" + filename

			_, errI := ctrl.repoImage.Insert(int(res.ID), entities.Images{RoomsID: res.ID, Image: newImage.Image})
			if errI != nil {
				return c.JSON(http.StatusInternalServerError, common.InternalServerError(nil, "error in upload image", nil))
			}
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create Room", ToRoomCreateResponseFormat(res)))
	}
}

func (ctrl *RoomsController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ctrl.repo.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Room", ToRoomGetResponseFormat(res)))
	}
}

func (ctrl *RoomsController) GetByUID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenId(c)
		res, err := ctrl.repo.GetByUID(uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Room", ToRoomGetResponseFormat(res)))
	}
}

func (ctrl *RoomsController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		Id, _ := strconv.Atoi(c.Param("roomid"))

		res, err := ctrl.repo.GetById(uint(Id))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Room By ID", ToRoomByIdGetResponseFormat(res)))
	}
}

func (ctrl *RoomsController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		roomId, _ := strconv.Atoi(c.Param("roomid"))
		UserID := middlewares.ExtractTokenId(c)

		var UpdateRoom = UpdateRoomRequestFormat{}

		if err := c.Bind(&UpdateRoom); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ctrl.repo.Update(uint(roomId), uint(UserID), UpdateRoom.ToUpdateRoomRequestFormat())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "success update room", ToUpdateRoomResponseFormat(res)))
	}
}

func (ctrl *RoomsController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenId(c)
		roomId, _ := strconv.Atoi(c.Param("roomid"))
		err := ctrl.repo.Delete(uint(roomId), uint(UserID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "success delete category", err))
	}
}
