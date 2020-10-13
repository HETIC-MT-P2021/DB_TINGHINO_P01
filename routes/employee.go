package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/HETIC-MT-P2021/DB_TINGHINO_P01/controllers"
)

// AddEmployeeRoutes define all employee routes.
func AddEmployeeRoutes(e *echo.Echo) {
	e.GET("/employee/:id", controllers.GetEmployee)
	e.GET("/employee", controllers.GetAllEmployee)
}

