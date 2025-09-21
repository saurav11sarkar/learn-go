package main

import "fmt"

// func printFunction(arr *[3]int) {
// 	fmt.Println(arr)
// }

// func sliceLiteral() {
// 	a := []int{1, 2, 3}
// 	fmt.Println(a)
// 	fmt.Println(len(a))
// 	fmt.Println(cap(a))

// 	fmt.Println("selice :", a, "len", len(a), "cap", cap(a))
// }

// func buildSlice() {
// 	// s := make([]int, 5)
// 	s := make([]int, 5, 10)
// 	s[4] = 4
// 	fmt.Println("selice :", s, "len", len(s), "cap", cap(s))
// }

func emptySlice() {
	s := []int{}
	s = append(s, 2,3)
	fmt.Println("selice :", s, "len", len(s), "cap", cap(s))
}

func main() {
	// a := 10

	// p := &a
	// *p = 20

	// fmt.Println("Hello, World!", a)

	// arr := [3]int{1, 2, 3}
	// printFunction(&arr)
	// sliceLiteral()
	// buildSlice()

	emptySlice()
}
