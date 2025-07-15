package main

// func add(a int, b int) int {
// 	return a + b
// }

// func getLanguages() (string, string, int){
// 	return "Go", "Python", 22
// }

func processit() func(a int) int {
	return func(a int) int {
		return a + 10
	}
}

func main() {
	// result := add(5, 3)
	// println("The sum is:", result)

	// language1, language2, val3 := getLanguages()
	// println("Languages:", language1, language2, val3)

	result := processit()
	println("Processed value:", result(5))
}