package models

import (
	"database/sql"
	"errors"

	"github.com/HETIC-MT-P2021/DB_TINGHINO_P01/db"

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

// GetOrder returns data of a order.
func GetOrder(orderNumber int) (Order, error) {
	var order Order

	const query = `SELECT * FROM orders WHERE orderNumber = ?`
	err := db.DB.
		QueryRow(query, orderNumber).
		Scan(
			&order.OrderNumber,
			&order.OrderDate,
			&order.RequiredDate,
			&order.ShippedDate,
			&order.Status,
			&order.Comments,
			&order.CustomerNumber,
		)

	if err == sql.ErrNoRows {
		return order, errors.New("order is not found")
	}

	if err != nil {
		return order, err
	}

	return order, nil
}

// GetAllOrder returns list of data of all orders.
func GetAllOrder() ([]Order, error) {
	var orderList []Order

	const query = `SELECT
						orderNumber,
						orderDate,
						requiredDate,
						shippedDate,
						status,
						comments,
						customerNumber
					FROM orders
					ORDER BY orderNumber`

	rows, err := db.DB.Query(query)

	if err != nil {
		return orderList, err
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		err = rows.
			Scan(
				&order.OrderNumber,
				&order.OrderDate,
				&order.RequiredDate,
				&order.ShippedDate,
				&order.Status,
				&order.Comments,
				&order.CustomerNumber,
			)

		if err != nil {
			return orderList, err
		}

		orderList = append(orderList, order)
	}

	err = rows.Err()

	if err != nil {
		return orderList, err
	}

	return orderList, nil
}

// GetOrderDetailWithProducts returns the list of products of an order with these details.
func GetProductsByOrderNumber(orderNumber int) ([]Product, error) {
	var productList []Product

	var orderDetail OrderDetail

	const query = `SELECT
       					p.productCode,
       					p.productName,
       					p.productLine,
       					p.productScale,
       					p.productVendor,
       					p.productDescription,
       					p.quantityInStock,
       					p.buyPrice,
       					p.MSRP,
       					od.orderNumber,
       					od.productCode,
       					od.quantityOrdered,
       					od.orderLineNumber,
       					od.priceEach
					FROM ((orders as o
							JOIN orderdetails as od
							ON o.orderNumber = od.orderNumber)
						JOIN products as p
						ON p.productCode = od.productCode)
					WHERE o.orderNumber = ?
					GROUP BY od.orderNumber, p.productCode`

	rows, err := db.DB.Query(query, orderNumber)

	if err != nil {
		return productList, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err = rows.
			Scan(
				&product.ProductCode,
				&product.ProductName,
				&product.ProductLine,
				&product.ProductScale,
				&product.ProductVendor,
				&product.ProductDescription,
				&product.QuantityInStock,
				&product.BuyPrice,
				&product.MSRP,
				&orderDetail.OrderNumber,
				&orderDetail.ProductCode,
				&orderDetail.QuantityOrdered,
				&orderDetail.PriceEach,
				&orderDetail.OrderLineNumber,
			)

		od := orderDetail

		product.OrderDetail = &od

		if err != nil {
			return productList, err
		}

		productList = append(productList, product)
	}

	err = rows.Err()

	if err != nil {
		return productList, err
	}

	return productList, nil
}
