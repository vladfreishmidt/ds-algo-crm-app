package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePass := "password"
	databaseName := "crm_db"
	database, error := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if error != nil {
		panic(error.Error())
	}
	return database
}

func GetCustomerById(customerId int) Customer {
	var database *sql.DB = GetConnection()

	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM customers WHERE ID=?", customerId)
	if error != nil {
		panic(error.Error())
	}
	//fmt.Println(rows)
	customer := Customer{}

	for rows.Next() {
		var customerId int
		var customerName string
		var SSN string
		error = rows.Scan(&customerId, &customerName, &SSN)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = SSN
	}
	defer database.Close()
	return customer
}
