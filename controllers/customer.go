package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"github.com/HETIC-MT-P2021/DB_TINGHINO_P01/models"
)

// GetCustomer returns a JSON object for one customer.
func GetCustomer(c echo.Context) error {
	customerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	res, err := models.GetCustomerByID(customerID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, http.StatusText(http.StatusOK), res))
}

// GetAllCustomer returns a JSON list of customer.
func GetAllCustomer(c echo.Context) error {
	res, err := models.GetAllCustomer()
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	if len(res) == 0 {
		return c.JSON(http.StatusOK, SetResponse(http.StatusOK, "product is empty", EmptyValue))
	}

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, http.StatusText(http.StatusOK), res))
}