package main

import (
	config "app_airbnb/configs"
	_adminController "app_airbnb/delivery/controllers/admin"
	_authController "app_airbnb/delivery/controllers/auth"

	_categoryController "app_airbnb/delivery/controllers/categories"
	_imagesController "app_airbnb/delivery/controllers/images"
	_roomsController "app_airbnb/delivery/controllers/rooms"
	_userController "app_airbnb/delivery/controllers/user"

	"app_airbnb/delivery/route"
	_adminRepo "app_airbnb/repository/admin"
	_authRepo "app_airbnb/repository/auth"

	_categoryRepo "app_airbnb/repository/categories"
	_imagesRepo "app_airbnb/repository/images"
	_roomsRepo "app_airbnb/repository/rooms"
	_userRepo "app_airbnb/repository/user"

	awss3 "app_airbnb/utils/aws_S3"
	// midtrans "app_airbnb/utils/midtrans"
	utils "app_airbnb/utils/mysql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()
	db := utils.InitDB(config)
	awsConn := awss3.InitS3(config.S3_KEY, config.S3_SECRET, config.S3_REGION)

	//REPOSITORY-DATABASE
	userRepo := _userRepo.New(db)
	authRepo := _authRepo.New(db)
	adminRepo := _adminRepo.New(db)
	roomsRepo := _roomsRepo.New(db)
	categoryRepo := _categoryRepo.New(db)
	imageRepo := _imagesRepo.New(db)

	//CONTROLLER
	userController := _userController.New(userRepo)
	authController := _authController.New(authRepo)
	adminController := _adminController.New(adminRepo)
	roomsController := _roomsController.New(roomsRepo)
	categoryController := _categoryController.New(categoryRepo)
	imageController := _imagesController.New(imageRepo, awsConn)

	e := echo.New()

	route.RegisterPath(e, userController, adminController, authController, categoryController, roomsController, imageController)

	// c := midtrans.InitConnection()
	// midtrans.CreateTransaction(c)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
