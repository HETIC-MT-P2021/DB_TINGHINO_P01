package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// InitRoutes initiates all routes.
func InitRoutes(e *echo.Echo) {
	AddCustomerRoutes(e)
	AddOrderRoutes(e)
	AddEmployeeRoutes(e)
	AddOfficeRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, e.Routes())
	})
}
