package main

import (
	"fmt"
	"maps"
)

func main() {
	// m := make(map[string]string)
	// m["name"] = "nazeer"
	// m["age"] = "22"
	// fmt.Println(m["name"], m["age"])

	// m := make(map[string]int)
	// m["nazeer"] = 22
	// m["age"] = 23
	// fmt.Println(m["nazeer"], m["age"])
	// fmt.Println(len(m))

	// fmt.Println(m)
	// delete(m, "nazeer")
	// fmt.Println(m)

	// m := map[string]int{"Price": 40, "Quantity": 100}
	// // fmt.Println(m)
	// v,ok := m["Price"]
	// if ok {
	// 	fmt.Println("Price exists:", v)
	// } else {
	// 	fmt.Println("Price does not exist")
	// }

	m1 := map[string]int{"Price": 40, "Quantity": 100}
	m2 := map[string]int{"Price": 40, "Quantity": 100}
	fmt.Println(maps.Equal(m1, m2))
}
