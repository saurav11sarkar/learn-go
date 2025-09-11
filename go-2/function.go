package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func main() {
	result := add(10, 20)
	fmt.Println("The sum is:", result)
}
