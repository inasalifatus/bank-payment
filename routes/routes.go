package routes

import (
	"github.com/inasalifatus/bank-payment/constants"
	"github.com/inasalifatus/bank-payment/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	// ----- to login and logout
	e.POST("/login", controllers.LoginCustomerController)
	eJwt.PUT("/customers/:id/logout", controllers.LogoutCustomerController)
	e.POST("/customers", controllers.CreateCustomersController)
	e.GET("/customers/:id", controllers.GetOneCustomersByIdController)
	e.DELETE("/customers/:id", controllers.DeleteCustomersByIdController)
	e.PUT("/customers/:id", controllers.UpdateCustomersController)
}
