package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/HETIC-MT-P2021/DB_TINGHINO_P01/controllers"
)

// AddCustomerRoutes define all customer routes.
func AddCustomerRoutes(e *echo.Echo) {
	e.GET("/customer/:id", controllers.GetCustomer)
	e.GET("/customer", controllers.GetAllCustomer)
}

