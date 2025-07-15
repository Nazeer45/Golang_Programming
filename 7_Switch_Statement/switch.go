package main

import "fmt"

func main() {
	// Example of a switch statement
	day := "Sunday"

	switch day {
	case "Monday":
		fmt.Println("It's the start of the week!")
	case "Friday":
		fmt.Println("It's almost the weekend!")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a regular weekday.")
	}

	//type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("I am an int:", i)
		case string:
			fmt.Println("I am a string:", i)
		case bool:
			fmt.Println("I am a bool:", i)
		default:
			fmt.Println("I am of a different type:", t)
		}
	}

		whoAmI(42)
		whoAmI("Hello")
}
