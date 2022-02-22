package route

import (
	"app_airbnb/delivery/controllers/address"
	"app_airbnb/delivery/controllers/auth"
	"app_airbnb/delivery/controllers/admin"
	"app_airbnb/delivery/controllers/user"
	"app_airbnb/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc *user.UserController, a *admin.AdminController, ac *address.AddressController, aa *auth.AuthController) {
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
	e.GET("users/me", uc.GetById(), middlewares.JwtMiddleware())
	e.PUT("users/me", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("users/me", uc.Delete(), middlewares.JwtMiddleware())
	
	//ROUTE Admin
	e.POST("admin/register", a.Register())
	e.GET("admin/users", a.GetAll(), middlewares.JwtMiddleware())

	//ROUTE ADDRESS
	ea := e.Group("")
	ea.POST("address", ac.Register(), middlewares.JwtMiddleware())
	ea.GET("address", ac.Get())
	ea.GET("address/:id", ac.GetById(), middlewares.JwtMiddleware())
	ea.PUT("address/:id", ac.Update(), middlewares.JwtMiddleware())
	ea.DELETE("address/:id", ac.Delete(), middlewares.JwtMiddleware())
}
