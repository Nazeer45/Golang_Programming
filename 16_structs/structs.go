package main

import (
	"fmt"
	"time"
)

type order struct {
	id        int
	amount    float32
	status    string
	createdAt time.Time
}

// constructor
// func newOrder(id int, amount float32, status string) *order {
// 	myOrder := order{
// 		id:        id,
// 		amount:    amount,
// 		status:    status,
// 	}
// 	return &myOrder
// }

//setter
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// getter
// func (o order) getStatus() float32 {
// 	return o.amount
// }

func main() {
	// myOrder:= order{
	// 	id:        1,
	// 	amount:    99.00,
	// 	status:    "Received",
	// 	createdAt: time.Now(),
	// }
	// fmt.Println(myOrder)
	// fmt.Println("amount:", myOrder.amount)
	// myOrder.status="Paid"
	// fmt.Println("status:", myOrder.status)
	
	// myOrder.changeStatus("Confirmed")
	// fmt.Println(myOrder)

	// fmt.Println(myOrder.getStatus())

	// myOrder2 := newOrder(2, 199.99, "Pending")
	// fmt.Println(myOrder2)

	language := struct {
		name    string
		isGood bool
	}{"golang", true}
	fmt.Println(language)
	 
}