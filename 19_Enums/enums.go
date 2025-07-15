package main

import "fmt"

// integer types
// type OrderStatus int
// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

// string types
type OrderStatus string
const (
	Received OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Prepared OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to:", status)
}

func main() {
	changeOrderStatus(Received)
	changeOrderStatus(Confirmed)
	changeOrderStatus(Prepared)
	changeOrderStatus(Delivered)
}