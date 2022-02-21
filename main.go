package main

import (
	config "app_airbnb/configs"
	_addressController "app_airbnb/delivery/controllers/address"
	_authController "app_airbnb/delivery/controllers/auth"
	_userController "app_airbnb/delivery/controllers/user"
	_ownerController "app_airbnb/delivery/controllers/owner"

	"app_airbnb/delivery/route"
	_addressRepo "app_airbnb/repository/address"
	_authRepo "app_airbnb/repository/auth"
	_userRepo "app_airbnb/repository/user"
	_ownerRepo "app_airbnb/repository/owner"

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
	ownerRepo := _ownerRepo.New(db)

	//CONTROLLER
	userController := _userController.New(userRepo)
	addressController := _addressController.New(addressRepo)
	authController := _authController.New(authRepo)
	ownerController := _ownerController.New(ownerRepo)


	e := echo.New()

	route.RegisterPath(e, userController, ownerController, addressController, authController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
