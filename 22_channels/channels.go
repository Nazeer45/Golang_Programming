package main

import (
	"fmt"
)

// func processNum(numChan chan int){
// 	fmt.Println("Processing number:", <- numChan)
// }

//sending
// func sum(result chan int, a int, b int) {
// 	numResult := a + b
// 	result <- numResult
// }

//Buffered channels
// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("Sending email to:", email)
// 		time.Sleep(time.Second) // Simulate time taken to send an email
// 	}
// }

func main() {
	// This is a simple example of using channels in Go, both sending and receiving messages are happening in single goroutine causing a deadlock.
	// messageChan := make(chan string)
	// messageChan <- "Hello, World!"
	// msg := <-messageChan
	// fmt.Println(msg)

	// Another example of using channels
	// numChan1 := make(chan int)
	// go processNum(numChan1)
	// numChan1 <- 42
	// time.Sleep(time.Second * 2)

	// Example of receiving messages from a channel in a separate goroutine
	// result := make(chan int)
	// go sum(result, 4, 5)
	// res := <-result
	// fmt.Println("Sum:", res)

	// Example using buffered channels
	// emailChan := make(chan string, 100)
	// done := make(chan bool)
	// go emailSender(emailChan, done)
	// for i := 0; i < 10; i++ {
	// 	emailChan <- fmt.Sprintf("%d@example.com", i)
	// }
	// fmt.Println("All emails sent to the channel")
	// // Close the channel to signal that no more emails will be sent
	// close(emailChan)
	// <-done
	// fmt.Println("All emails processed")

	//Example of listening from multiple channels
	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
			chan1 <- 10
	}()

	go func() {
			chan2 <- "pong"
	}()

	for i:= 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received from chan1:", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Received from chan2:", chan2Val)
		}
	}
}