package main

import (
	"fmt"

	"example.go/mathlid"
)

func main() {
	fmt.Println("Hello, World!")
	var a, b int
	fmt.Print("Enter your Numbers 1 : ")
	fmt.Scanln(&a)
	fmt.Print("Enter your Numbers 2 : ")
	fmt.Scanln(&b)
	// add(a, b)
	mathlid.Multiply(a, b)
}
