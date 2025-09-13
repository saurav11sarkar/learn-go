package main

import "fmt"

var a = 10
var b = 20

func main() {
	fmt.Println("Hello, init function!")
	fmt.Println("a + b =", a+b)

	func(x, y int) {
		fmt.Println("Anonymous function called!", x, y)
	}(a, b)
}

func init() {
	fmt.Println("Init function called!", a, b)
	a = 100
	b = 200
}

func init() {
	fmt.Println("Init function called again!", a, b)
}
