package main

import "fmt"

func main() {
	// nums := []int{2, 6, 9, 7}

	// for i, nums := range nums {
	// 	fmt.Printf("Index: %d, Value: %d\n", i, nums)
	// }

	m:= map[string]string{
		"firstName :":  "Shaik",
		"lastName :":   "Nazeer",
		"occupation :": "Software Engineer",
		"age":         "30",}

		for key, value := range m {
			fmt.Println(key, value)
		}
}