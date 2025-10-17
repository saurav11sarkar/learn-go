package main

import "fmt"

// func a() {
// 	a := 1
// 	fmt.Println("before defer", a)
// 	defer fmt.Println("first def", a)
// 	a++
// 	fmt.Println("after defer", a)
// 	defer fmt.Println("second defer", a)

// }

// func name(a int,b int) (result int){
// 	result = a + b
// 	return
// }

func calcute() (result int) {
	fmt.Println("calcute start", result)
	show := func() {
		result = result + 10
		fmt.Println("result:", result)
	}

	defer show()

	result = 5
	fmt.Println("calcute end", result)
	return
}

func cal() int {
	result := 5
	show := func() {
		result = result + 10
		fmt.Println("result:", result)
	}
	defer show()
	result = 5
	fmt.Println("main end", result)
	return result
}

func main() {
	// a()

	// fmt.Println(name(1, 2))
	returning := calcute()
	fmt.Println("returning:", returning)
	noreturning := cal()
	fmt.Println("noreturning:", noreturning)
}
