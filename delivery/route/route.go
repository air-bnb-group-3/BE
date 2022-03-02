package route

import (
	"app_airbnb/delivery/controllers/admin"
	"app_airbnb/delivery/controllers/auth"

	"app_airbnb/delivery/controllers/booking"
	"app_airbnb/delivery/controllers/categories"
	"app_airbnb/delivery/controllers/images"
	"app_airbnb/delivery/controllers/rooms"
	"app_airbnb/delivery/controllers/user"
	"app_airbnb/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo,
	uc *user.UserController,
	a *admin.AdminController,
	aa *auth.AuthController,
	cat *categories.CategoriesController,
	r *rooms.RoomsController,
	i *images.ImagesController,
	b *booking.BookingController,
) {

	//CORS
	e.Use(middleware.CORS())

	//LOGGER
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	//ROUTE REGISTER - LOGIN USERS
	e.POST("users/register", uc.Register())
	e.POST("users/login", aa.Login())

	//ROUTE USERS
	e.GET("/users", uc.GetById(), middlewares.JwtMiddleware())
	e.PUT("/users", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users", uc.Delete(), middlewares.JwtMiddleware())

	//ROUTE Admin
	e.POST("admin/register", a.Register())
	e.POST("admin/login", aa.AdminLogin())
	e.GET("admin/users", a.GetAll(), middlewares.JwtMiddleware())

	// ROUTE Category
	e.POST("/categories", cat.Insert(), middlewares.JwtMiddleware())
	e.GET("/categories", cat.GetAll())
	e.GET("/categories/:categoryid", cat.GetById())
	e.PUT("/categories/:categoryid", cat.Update(), middlewares.JwtMiddleware())
	e.DELETE("/categories/:categoryid", cat.Delete(), middlewares.JwtMiddleware())

	//ROUTE Rooms
	e.POST("/rooms", r.Insert(), middlewares.JwtMiddleware())
	e.GET("/rooms/all", r.GetAll())
	e.GET("/rooms", r.GetByUID(), middlewares.JwtMiddleware())
	e.GET("/rooms/:roomid", r.GetById())
	e.PUT("/rooms/:roomid", r.Update(), middlewares.JwtMiddleware())
	e.DELETE("/rooms/:roomid", r.Delete(), middlewares.JwtMiddleware())

	//ROUTE Images
	e.POST("/images", i.Insert(), middlewares.JwtMiddleware())

	//ROUTE Booking
	e.POST("/booking", b.Create(), middlewares.JwtMiddleware())
	e.POST("/booking/:booking_id", b.CreatePayment(), middlewares.JwtMiddleware())
	e.POST("/booking/payment/callback", b.CallBack())

	e.GET("/booking", b.GetByUserID(), middlewares.JwtMiddleware())
	e.GET("/booking/:booking_id", b.GetById(), middlewares.JwtMiddleware())
	e.PUT("/booking", b.Update(), middlewares.JwtMiddleware())

}
