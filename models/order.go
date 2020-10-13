package models

import (
	"gopkg.in/guregu/null.v4"
)

type Order struct {
	OrderNumber    int         `json:"orderNumber"`
	OrderDate      string      `json:"orderDate"`
	RequiredDate   string      `json:"requiredDate"`
	ShippedDate    null.String `json:"shippedDate"`
	Status         string      `json:"status"`
	Comments       null.String `json:"comments"`
	CustomerNumber int         `json:"customerNumber"`
}

type OrderDetail struct {
	OrderNumber     int      `json:"orderNumber"`
	ProductCode     string   `json:"productCode"`
	QuantityOrdered string   `json:"quantityOrdered"`
	PriceEach       string   `json:"priceEach"`
	OrderLineNumber string   `json:"orderLineNumber"`
	Product         *Product `json:"product"`
}
