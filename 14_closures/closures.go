package main

import "fmt"

func counter() func() int {
	var count int =0
	return func() int{
		count++
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2
}