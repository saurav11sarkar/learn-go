package main

import "fmt"

func main() {
	println("Hello, World!")

	var a, b int

	print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	fmt.Println("Sum: ", a+b)
}
