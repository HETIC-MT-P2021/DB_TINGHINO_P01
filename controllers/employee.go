package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"github.com/HETIC-MT-P2021/DB_TINGHINO_P01/models"
)

// GetEmployee returns a JSON object for one employee.
func GetEmployee(c echo.Context) error {
	employeeNumber, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	res, err := models.GetEmployeeByNumber(employeeNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, http.StatusText(http.StatusOK), res))
}

// GetAllEmployee returns a JSON list of employee.
func GetAllEmployee(c echo.Context) error {
	res, err := models.GetAllEmployee()
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	if len(res) == 0 {
		return c.JSON(http.StatusOK, SetResponse(http.StatusOK, "product is empty", EmptyValue))
	}

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, http.StatusText(http.StatusOK), res))
}
