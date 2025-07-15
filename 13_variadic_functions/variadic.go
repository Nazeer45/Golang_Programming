package main

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main(){
	result := sum(2,4,9,7,8)
	println("The sum is:", result)
}