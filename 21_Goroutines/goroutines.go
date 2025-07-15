package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done() // Decrement the counter when the goroutine completes
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup 
	for i:=0; i<10; i++{
		wg.Add(1) // Increment the counter for each task
		go task(i, &wg) // Start a goroutine for each task
	}
	wg.Wait() // Wait for all goroutines to finish
}
