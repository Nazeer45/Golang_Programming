package main

import "fmt"

func main() {
	// while loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	//infinite loop
	for {
		fmt.Println("Infinite Loop")
		break // to exit the infinite loop
	}

	// classic for loop
	for j := 1; j <= 5; j++ {
		if j == 2 {
			continue // skip the rest of the loop when j is 2
		}
		fmt.Println(j)
	}

	// range based for loop
	for k := range 3 {
		fmt.Println("Range based loop iteration:", k)
	}


}