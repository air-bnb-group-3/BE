package rooms

import (
	"app_airbnb/delivery/controllers/common"
	"app_airbnb/delivery/middlewares"
	"app_airbnb/repository/rooms"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoomsController struct {
	repo rooms.Rooms
}

func New(repository rooms.Rooms) *RoomsController {
	return &RoomsController{
		repo: repository,
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

func (ctrl *RoomsController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		roomId, _ := strconv.Atoi(c.Param("roomid"))

		res, err := ctrl.repo.GetById(uint(roomId))

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
