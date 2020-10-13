package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/HETIC-MT-P2021/DB_TINGHINO_P01/controllers"
)

// AddOfficeRoutes define all office routes.
func AddOfficeRoutes(e *echo.Echo) {
	e.GET("/office/:id/employee", controllers.GetEmployeesByOfficeCode)
}
