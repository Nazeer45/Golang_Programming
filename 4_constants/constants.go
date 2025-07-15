package main

import "fmt"

const age = 32
var animal = "Dog"
// bird := "Parrot"  // This line will cause an error because 'bird' is not declared with 'var' or 'const' outside a function

func main() {
	const name = "Shaik Nazeer"
	// const age = 23

	// fmt.Println(age)

	const (
		port = 8080
		host = "localhost"
	)

	fmt.Println(port, host)

}