package main

import (
	config "app_airbnb/configs"
	_adminController "app_airbnb/delivery/controllers/admin"
	_authController "app_airbnb/delivery/controllers/auth"

	_categoryController "app_airbnb/delivery/controllers/categories"
	_roomsController "app_airbnb/delivery/controllers/rooms"
	_userController "app_airbnb/delivery/controllers/user"

	"app_airbnb/delivery/route"
	_adminRepo "app_airbnb/repository/admin"
	_authRepo "app_airbnb/repository/auth"

	_categoryRepo "app_airbnb/repository/categories"
	_roomsRepo "app_airbnb/repository/rooms"
	_userRepo "app_airbnb/repository/user"

	utils "app_airbnb/utils/mysql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()
	db := utils.InitDB(config)

	//REPOSITORY-DATABASE
	userRepo := _userRepo.New(db)
	authRepo := _authRepo.New(db)
	adminRepo := _adminRepo.New(db)
	roomsRepo := _roomsRepo.New(db)
	categoryRepo := _categoryRepo.New(db)

	//CONTROLLER
	userController := _userController.New(userRepo)
	authController := _authController.New(authRepo)
	adminController := _adminController.New(adminRepo)
	roomsController := _roomsController.New(roomsRepo)
	categoryController := _categoryController.New(categoryRepo)

	e := echo.New()

	route.RegisterPath(e, userController, adminController, authController, categoryController, roomsController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
