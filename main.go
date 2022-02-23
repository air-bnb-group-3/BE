package main

import (
	config "app_airbnb/configs"
	_adminController "app_airbnb/delivery/controllers/admin"
	_authController "app_airbnb/delivery/controllers/auth"
	_userController "app_airbnb/delivery/controllers/user"

	"app_airbnb/delivery/route"
	_adminRepo "app_airbnb/repository/admin"
	_authRepo "app_airbnb/repository/auth"
	_userRepo "app_airbnb/repository/user"

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
	authRepo := _authRepo.New(db)
	adminRepo := _adminRepo.New(db)

	//CONTROLLER
	userController := _userController.New(userRepo)
	authController := _authController.New(authRepo)
	adminController := _adminController.New(adminRepo)

	e := echo.New()

	route.RegisterPath(e, userController, adminController, authController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
