package main

import (
	config "app_airbnb/configs"
	_addressController "app_airbnb/delivery/controllers/address"
	_authController "app_airbnb/delivery/controllers/auth"
	_userController "app_airbnb/delivery/controllers/user"
	_adminController "app_airbnb/delivery/controllers/admin"

	"app_airbnb/delivery/route"
	_addressRepo "app_airbnb/repository/address"
	_authRepo "app_airbnb/repository/auth"
	_userRepo "app_airbnb/repository/user"
	_adminRepo "app_airbnb/repository/admin"

	"app_airbnb/utils"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()
	db := utils.InitDB(config)

	//REPOSITORY-DATABASE
	userRepo := _userRepo.New(db)
	addressRepo := _addressRepo.New(db)
	authRepo := _authRepo.New(db)
	adminRepo := _adminRepo.New(db)

	//CONTROLLER
	userController := _userController.New(userRepo)
	addressController := _addressController.New(addressRepo)
	authController := _authController.New(authRepo)
	adminController := _adminController.New(adminRepo)


	e := echo.New()

	route.RegisterPath(e, userController, adminController, addressController, authController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
