package models

import (
	"github.com/HETIC-MT-P2021/DB_TINGHINO_P01/db"
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

// GetEmployeesByOfficeCode returns the list of employees of an office.
func GetEmployeesByOfficeCode(officeCode int) ([]Employee, error) {
	var employeeList []Employee

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
					WHERE e.officeCode = ?
					ORDER BY e.employeeNumber;`

	rows, err := db.DB.Query(query, officeCode)

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
