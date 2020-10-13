package models

import (
	"gopkg.in/guregu/null.v4"
)

type Customer struct {
	CustomerNumber              int         `json:"customerNumber"`
	CustomerName                string      `json:"customerName"`
	ContactLastName             string      `json:"contactLastName"`
	ContactFirstName            string      `json:"contactFirstName"`
	Phone                       string      `json:"phone"`
	AddressLine1                string      `json:"addressLine1"`
	AddressLine2                null.String `json:"addressLine2"`
	City                        string      `json:"city"`
	State                       null.String `json:"state"`
	PostalCode                  null.String `json:"postalCode"`
	Country                     string      `json:"country"`
	SalesRepEmployeeNumber      null.Int    `json:"salesRepEmployeeNumber"`
	CreditLimit                 float64     `json:"creditLimit"`
	Orders                      null.String `json:"orders"`
	TotalNumberOfProductOrdered null.Int    `json:"totalNumberOfProductOrdered"`
	TotalPrice                  null.Float  `json:"totalPrice"`
}
