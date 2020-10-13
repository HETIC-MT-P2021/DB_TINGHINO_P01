package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/HETIC-MT-P2021/DB_TINGHINO_P01/controllers"
)

// AddOrderRoutes define all order routes.
func AddOrderRoutes(e *echo.Echo) {
	e.GET("/order", controllers.GetAllOrder)
	e.GET("/order/:id", controllers.GetOrder)
	e.GET("/order/:id/product", controllers.GetProductsByOrderNumber)
}
