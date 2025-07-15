package main

import "fmt"

func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T any] struct {
	items []T
}

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// printSlice(nums)

	// names := []string{"goland", "python"}
	// printSlice(names)

	myStack := stack[string]{
		items: []string{"golang", "python", "java"},
	}
	fmt.Println("Stack items:", myStack.items)
}