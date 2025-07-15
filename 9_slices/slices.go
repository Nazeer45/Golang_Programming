package main

import (
	"fmt"
	"slices"
)

func main() {
	// var nums []int
	// fmt.Println(nums) // Output: []
	// fmt.Println(nums==nil)

	// var nums = make([]int, 2, 5)
	// fmt.Println(nums) // Output: [0 0]
	// fmt.Println(len(nums)) // Output: 2
	// fmt.Println(cap(nums)) // Output: 5

	// nums = append(nums, 1, 2)
	// fmt.Println(nums) // Output: [0 0 1 2]
	// fmt.Println(cap(nums)) // Output: 5

	// nums = append(nums, 3, 4,)
	// fmt.Println(nums) // Output: [0 0 1 2 3 4]
	// fmt.Println(cap(nums)) // Output: 5

	//copy function
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))
	// copy(nums2, nums)
	// fmt.Println(nums2) // Output: [2]

	//Equal function
	var nums1 = []int{1, 2, 3}
	var nums2 = []int{1, 2, 3}
	fmt.Println(slices.Equal(nums1, nums2)) // Output: false, slices cannot be compared directly

}
