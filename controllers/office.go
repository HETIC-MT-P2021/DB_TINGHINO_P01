package controllers

import (
	"net/http"
	"strconv"

	"github.com/HETIC-MT-P2021/DB_TINGHINO_P01/models"

	"github.com/labstack/echo/v4"
)

// GetProductsByOrderNumber returns a JSON object for all employees in an office.
func GetEmployeesByOfficeCode(c echo.Context) error {
	orderNumber, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	res, err := models.GetEmployeesByOfficeCode(orderNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, http.StatusText(http.StatusOK), res))
}
