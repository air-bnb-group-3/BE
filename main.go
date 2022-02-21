package main

import (
	config "app_airbnb/configs"
	addressController "app_airbnb/delivery/controllers/address"
	authController "app_airbnb/delivery/controllers/auth"
	userController "app_airbnb/delivery/controllers/user"
	"app_airbnb/delivery/route"
	addressRepo "app_airbnb/repository/address"
	authRepo "app_airbnb/repository/auth"
	userRepo "app_airbnb/repository/user"
	"app_airbnb/utils"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()
	db := utils.InitDB(config)

	//REPOSITORY-DATABASE
	userRepo := userRepo.New(db)
	addressRepo := addressRepo.New(db)
	authRepo := authRepo.New(db)

	//CONTROLLER
	userController := userController.New(userRepo)
	addressController := addressController.New(addressRepo)
	authController := authController.New(authRepo)

	e := echo.New()

	route.RegisterPath(
		e,
		userController,
		addressController,
		authController,
	)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
