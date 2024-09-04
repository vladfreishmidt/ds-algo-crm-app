package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var template_html = template.Must(template.ParseGlob("templates/*"))

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")
	template_html.ExecuteTemplate(w, "Home.tmpl", nil)
}

func View(writer http.ResponseWriter, request *http.Request) {
	//var customers []Customer
	//customers = GetCustomers()
	//log.Println(customers)
	var customerId int
	var customerIdStr string
	// customerIdStr = request.FormValue("id")
	customerIdStr = "2"
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)
	fmt.Println(customer)
	var customers []Customer
	customers = []Customer{customer}
	//  customers.append(customer)
	// template_html.ExecuteTemplate(writer, "View", customers)
	fmt.Println(customers)
}

func main() {
	fmt.Println("Server started on port :8080")

	http.HandleFunc("/", Home)
	http.HandleFunc("/view", View)

	http.ListenAndServe(":8080", nil)
}
