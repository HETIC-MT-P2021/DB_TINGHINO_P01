package controllers

import (
	"net/http"
	"strconv"

	"github.com/HETIC-MT-P2021/DB_TINGHINO_P01/models"

	"github.com/labstack/echo/v4"
)

// GetOrder returns a JSON object for one order.
func GetOrder(c echo.Context) error {
	orderNumber, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	res, err := models.GetOrder(orderNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, http.StatusText(http.StatusOK), res))
}

// GetAllOrder returns a JSON list of orders.
func GetAllOrder(c echo.Context) error {
	res, err := models.GetAllOrder()
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	if len(res) == 0 {
		return c.JSON(http.StatusOK, SetResponse(http.StatusOK, "order is empty", EmptyValue))
	}

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, http.StatusText(http.StatusOK), res))
}

// GetProductsByOrderNumber returns a JSON object for all product in an order.
func GetProductsByOrderNumber(c echo.Context) error {
	orderNumber, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	res, err := models.GetProductsByOrderNumber(orderNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, http.StatusText(http.StatusOK), res))
}
