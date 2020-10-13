package models

import (
	"gopkg.in/guregu/null.v4"
)

type Office struct {
	OfficeCode   string      `json:"officeCode"`
	City         string      `json:"city"`
	Phone        string      `json:"phone"`
	AddressLine1 string      `json:"addressLine1"`
	AddressLine2 null.String `json:"addressLine2"`
	State        null.String `json:"state"`
	Country      string      `json:"country"`
	PostalCode   string      `json:"postalCode"`
	Territory    string      `json:"territory"`
	Employee     *[]Employee `json:"employee"`
}
