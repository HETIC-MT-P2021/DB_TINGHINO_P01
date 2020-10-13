package models

import (
	"database/sql"
	"errors"
	"github.com/HETIC-MT-P2021/DB_TINGHINO_P01/db"

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

// GetCustomerByID returns data of a customer.
func GetCustomerByID(customerNumber int) (Customer, error) {
	var customer Customer

	const query = `SELECT
					    c.customerNumber,
					    c.customerName,
					    c.contactLastName,
					    c.contactFirstName,
					    c.phone,
					    c.addressLine1,
					    c.addressLine2,
					    c.city,
					    c.state,
					    c.postalCode,
					    c.country,
					    c.salesRepEmployeeNumber,
					    c.creditLimit,
					    GROUP_CONCAT(DISTINCT orders.orderNumber) as orders,
					    SUM(orders.quantityOrdered) as totalNumberOfProductOrdered,
					    SUM(orders.priceEach * orders.quantityOrdered) as totalPrice
					FROM customers AS c
					         LEFT JOIN
					     (SELECT o.customerNumber, od.orderNumber, od.productCode, od.priceEach, od.quantityOrdered
					      FROM orders as o
					               LEFT JOIN orderdetails AS od
					                         ON od.orderNumber = o.orderNumber
					      GROUP BY o.customerNumber, od.orderNumber, od.productCode
					     ) as orders
					     ON c.customerNumber = orders.customerNumber
					WHERE c.customerNumber = ?
					GROUP BY c.customerNumber;`

	err := db.DB.
		QueryRow(query, customerNumber).
		Scan(
			&customer.CustomerNumber,
			&customer.CustomerName,
			&customer.ContactLastName,
			&customer.ContactFirstName,
			&customer.Phone,
			&customer.AddressLine1,
			&customer.AddressLine2,
			&customer.City,
			&customer.State,
			&customer.PostalCode,
			&customer.Country,
			&customer.SalesRepEmployeeNumber,
			&customer.CreditLimit,
			&customer.Orders,
			&customer.TotalNumberOfProductOrdered,
			&customer.TotalPrice,
		)

	if err == sql.ErrNoRows {
		return customer, errors.New("customer is not found")
	}

	if err != nil {
		return customer, err
	}

	return customer, nil
}

// GetAllCustomer returns list of data of all customer.
func GetAllCustomer() ([]Customer, error) {
	var customerList []Customer

	const query = `SELECT
					    c.customerNumber,
					    c.customerName,
					    c.contactLastName,
					    c.contactFirstName,
					    c.phone,
					    c.addressLine1,
					    c.addressLine2,
					    c.city,
					    c.state,
					    c.postalCode,
					    c.country,
					    c.salesRepEmployeeNumber,
					    c.creditLimit,
					    GROUP_CONCAT(DISTINCT orders.orderNumber) as orders,
					    SUM(orders.quantityOrdered) as totalNumberOfProductOrdered,
					    SUM(orders.priceEach * orders.quantityOrdered) as totalPrice
					FROM customers AS c
					         LEFT JOIN
					     (SELECT o.customerNumber, od.orderNumber, od.productCode, od.priceEach, od.quantityOrdered
					      FROM orders as o
					               LEFT JOIN orderdetails AS od
					                         ON od.orderNumber = o.orderNumber
					      GROUP BY o.customerNumber, od.orderNumber, od.productCode
					     ) as orders
					     ON c.customerNumber = orders.customerNumber
					GROUP BY c.customerNumber;`

	rows, err := db.DB.Query(query)

	if err != nil {
		return customerList, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer Customer
		err = rows.
			Scan(
				&customer.CustomerNumber,
				&customer.CustomerName,
				&customer.ContactLastName,
				&customer.ContactFirstName,
				&customer.Phone,
				&customer.AddressLine1,
				&customer.AddressLine2,
				&customer.City,
				&customer.State,
				&customer.PostalCode,
				&customer.Country,
				&customer.SalesRepEmployeeNumber,
				&customer.CreditLimit,
				&customer.Orders,
				&customer.TotalNumberOfProductOrdered,
				&customer.TotalPrice,
			)

		if err != nil {
			return customerList, err
		}

		customerList = append(customerList, customer)
	}

	err = rows.Err()

	if err != nil {
		return customerList, err
	}

	return customerList, nil
}
