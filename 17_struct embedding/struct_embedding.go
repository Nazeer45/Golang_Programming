package main

import (
	"fmt"
)

type customer struct {
	name  string
	email string
	age   int
}

type order struct {
	id        int
	amount    float32
	status    string
	customer  // embedding customer struct
}

func main() {

	// newCustomer := customer{
	// 	name:  "John Doe",
	// 	email: "hjvjj",
	// 	age:   30,
	// }

	myOrder := order{
		id:       1,
		amount:   99.00,
		status:   "Received",
		// customer: newCustomer,
		customer: customer{
			name:  "John Doe",
			email: "hjvjj",
			age:   30,
		},
	}
	fmt.Println(myOrder)
	myOrder.customer.name = "Shaik"
	fmt.Println(myOrder)
}
