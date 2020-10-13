# SQL JOIN

## Usage

Required development environment:
- [Docker](https://www.docker.com)
- [Docker Compose](https://docs.docker.com/compose/install/)

Configure the development environment on your local machine:
```bash
$ git clone https://github.com/HETIC-MT-P2021/DB_TINGHINO_P01.git
$ cd DB_TINGHINO_P01
$ make compose:up
```

You can now access the api: [http://localhost:1323/](http://localhost:1323/).


## Endpoints


```bash
## Get all customers and infos (allOrders, totalNumberOfProductOrdered, totalPrice, )
GET: /customer/:id

## Get customer by id
GET: /customer/:id

-——

## Get all products
GET: /order

## Get product by id
GET: /order/:id

## Get all products of an order
GET: /order/:id/product

-——

## Get all employees (include office relationship)
GET: /employee

## Get employee by id
GET: /employee/:id

-——

## Get all employees of an office (include office relationship)
GET /office/:id/employee
```
