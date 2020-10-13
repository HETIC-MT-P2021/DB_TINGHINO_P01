package models

import (
	"gopkg.in/guregu/null.v4"
)

type Employee struct {
	EmployeeNumber int      `json:"employeeNumber"`
	LastName       string   `json:"lastName"`
	FirstName      string   `json:"firstName"`
	Extension      string   `json:"extension"`
	Email          string   `json:"email"`
	OfficeCode     string   `json:"officeCode"`
	ReportsTo      null.Int `json:"reportsTo"`
	JobTitle       string   `json:"jobTitle"`
	Office         *Office  `json:"office"`
}
