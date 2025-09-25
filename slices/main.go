package main

import "fmt"

func main() {
	// var nums []int
	// nums = append(nums, 1)

	// fmt.Println(nums)
	// if nums == nil {
	// 	fmt.Println("nil!")
	// }

	var nums = make([]int, 2, 5)
	fmt.Println(nums, len(nums), cap(nums))
}
