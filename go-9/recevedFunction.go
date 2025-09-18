package main

import "fmt"

func recevedFunction(x int) {
	fmt.Println("Inside recevedFunction, x =", x)
}

func main() {
	recevedFunction(10)
}
