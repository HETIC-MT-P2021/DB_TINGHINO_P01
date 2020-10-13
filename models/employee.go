package models

import (
	"database/sql"
	"errors"

	"github.com/HETIC-MT-P2021/DB_TINGHINO_P01/db"

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

// GetEmployeeByNumber returns data of a employee.
func GetEmployeeByNumber(employeeNumber int) (Employee, error) {
	var employee Employee
	var office Office

	const query = `SELECT
						e.employeeNumber,
						e.lastName,
						e.firstName,
						e.extension,
						e.email,
						e.officeCode,
						e.reportsTo,
						e.jobTitle,
						o.officeCode,
						o.city,
						o.phone,
						o.addressLine1,
						o.addressLine2,
						o.state,
						o.country,
						o.postalCode,
						o.territory
					FROM employees as e
					         JOIN offices as o
					             ON e.officeCode = o.officeCode
					WHERE e.employeeNumber = ?`

	err := db.DB.
		QueryRow(query, employeeNumber).
		Scan(
			&employee.EmployeeNumber,
			&employee.LastName,
			&employee.FirstName,
			&employee.Extension,
			&employee.Email,
			&employee.OfficeCode,
			&employee.ReportsTo,
			&employee.JobTitle,
			&office.OfficeCode,
			&office.City,
			&office.Phone,
			&office.AddressLine1,
			&office.AddressLine2,
			&office.State,
			&office.Country,
			&office.PostalCode,
			&office.Territory,
		)

	employee.Office = &office

	if err == sql.ErrNoRows {
		return employee, errors.New("employee is not found")
	}

	if err != nil {
		return employee, err
	}

	return employee, nil
}

// GetAllEmployee returns list of data of all employee.
func GetAllEmployee() ([]Employee, error) {
	var employeeList []Employee
	var office Office

	const query = `SELECT
						e.employeeNumber,
						e.lastName,
						e.firstName,
						e.extension,
						e.email,
						e.reportsTo,
						e.jobTitle,
						o.officeCode,
						o.city,
						o.phone,
						o.addressLine1,
						o.addressLine2,
						o.state,
						o.country,
						o.postalCode,
						o.territory
					FROM employees as e
					         JOIN offices as o
					             ON e.officeCode = o.officeCode
					ORDER BY e.employeeNumber`

	rows, err := db.DB.Query(query)

	if err != nil {
		return employeeList, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee Employee
		err = rows.
			Scan(
				&employee.EmployeeNumber,
				&employee.LastName,
				&employee.FirstName,
				&employee.Extension,
				&employee.Email,
				&employee.ReportsTo,
				&employee.JobTitle,
				&office.OfficeCode,
				&office.City,
				&office.Phone,
				&office.AddressLine1,
				&office.AddressLine2,
				&office.State,
				&office.Country,
				&office.PostalCode,
				&office.Territory,
			)

		o := office
		employee.Office = &o

		if err != nil {
			return employeeList, err
		}

		employeeList = append(employeeList, employee)
	}

	err = rows.Err()

	if err != nil {
		return employeeList, err
	}

	return employeeList, nil
}
